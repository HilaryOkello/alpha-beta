package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"alpha-beta/blockchain"
	"github.com/google/uuid"
)

func generateOrderID() string {
	return uuid.New().String()
}

func CreateHealthFacilityOrder(w http.ResponseWriter, r *http.Request) {
	var orderDetails struct {
		ManufacturerID   string `json:"manufacturer_id"`
		HealthFacilityID string `json:"health_facility_id"`
		VaccineDetails   string `json:"vaccine_details"`
		TransactionType  string `json:"transaction_type"`
		Status           string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&orderDetails); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not decode order details: %v", err)
		w.Write([]byte("could not decode order details"))
		return
	}

	detailsBytes, err := json.Marshal(orderDetails)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not marshal order details: %v", err)
		w.Write([]byte("could not marshal order details"))
		return
	}

	transaction := blockchain.VaccineTransaction{
		OrderID:   generateOrderID(),
		IsGenesis: false,
		Details:   string(detailsBytes),
	}

	blockchain.BlockChain.AddBlock(transaction)

	http.Redirect(w, r, "/inventory", http.StatusSeeOther)
}

func InventoryPage(w http.ResponseWriter, r *http.Request) {
	type InventoryData struct {
		PendingOrders   []*blockchain.Block
		FulfilledOrders []*blockchain.Block
	}

	var inventoryData InventoryData

	for _, block := range blockchain.BlockChain.Blocks {
		var details struct {
			Status string `json:"status"`
		}
		json.Unmarshal([]byte(block.Data.Details), &details)
		if details.Status == "Pending" {
			inventoryData.PendingOrders = append(inventoryData.PendingOrders, block)
		} else {
			inventoryData.FulfilledOrders = append(inventoryData.FulfilledOrders, block)
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/inventory.html"))
	if err := tmpl.Execute(w, inventoryData); err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}
