package handlers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"alpha-beta/blockchain"
)

// Vaccine represents a vaccine with details
type Vaccine struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	ExpiryDate   string `json:"expiry_date"`
	BatchNumber  string `json:"batch_number"`
}

// Serve the HTML form for the blockchain view
func BlockchainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}

// Serve the HTML form for creating a distributor order
func DistributorOrderPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/distributor-order.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}

// Serve the HTML form for creating a health facility order
func HealthFacilityOrderPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/pharmacy.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}

// Serve the HTML form for creating a new vaccine
func NewVaccinePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/manufacturer.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}

func CreateDistributorOrder(w http.ResponseWriter, r *http.Request) {
	var transaction blockchain.VaccineTransaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not create distributor order: %v", err)
		w.Write([]byte("could not create distributor order"))
		return
	}

	if transaction.TransactionType != "DistributorToManufacturer" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid transaction type for distributor order"))
		return
	}

	blockchain.BlockChain.AddBlock(transaction)
	resp, err := json.MarshalIndent(transaction, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not marshal payload: %v", err)
		w.Write([]byte("could not save distributor order"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func CreateHealthFacilityOrder(w http.ResponseWriter, r *http.Request) {
	var transaction blockchain.VaccineTransaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not create health facility order: %v", err)
		w.Write([]byte("could not create health facility order"))
		return
	}

	if transaction.TransactionType != "FacilityToDistributor" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid transaction type for health facility order"))
		return
	}

	blockchain.BlockChain.AddBlock(transaction)
	resp, err := json.MarshalIndent(transaction, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not marshal payload: %v", err)
		w.Write([]byte("could not save health facility order"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// newVaccine handles creating a new vaccine record
func NewVaccine(w http.ResponseWriter, r *http.Request) {
	var vaccine Vaccine
	if err := json.NewDecoder(r.Body).Decode(&vaccine); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not create: %v", err)
		w.Write([]byte("could not create new Vaccine"))
		return
	}

	h := sha256.New()
	io.WriteString(h, vaccine.BatchNumber+vaccine.ExpiryDate)
	vaccine.ID = fmt.Sprintf("%x", h.Sum(nil))

	resp, err := json.MarshalIndent(vaccine, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not marshal payload: %v", err)
		w.Write([]byte("could not save vaccine data"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
