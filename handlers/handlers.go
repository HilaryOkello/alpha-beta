package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"alpha-beta/blockchain"
)

type Manufacturer struct {
	ID   string
	Name string
}

var (
	availableVaccines      []Vaccine
	availableManufacturers []Manufacturer
	vaccineIDCounter       int
	generateUserID         int
)

// Vaccine represents a vaccine with details
type Vaccine struct {
	ID              string `json:"id"`
	Name            string `json:"type"`
	ManufacturerID  string `json:"manufacturer"`
	ManufactureDate string `json:"manufacture_date"`
	ExpiryDate      string `json:"expiry_date"`
	BatchNumber     string `json:"batch_number"`
	Quantity        int    `json: "quantity"`
}

// Serve the HTML form for the blockchain view
func IndexPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}

func SignupPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/signup.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}

func DistributorDashboard(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/distributor.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}

func HealthFacilityDashboard(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/pharmacy.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}

func ManufacturerDashboard(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/manufacturer.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}

func AddFacilityPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/add_facility.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}

func AddManufacturerPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/add_manufacturer.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}

// Serve the HTML form for creating a distributor order
func DistributorOrderPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/distributor_order.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}

// Serve the HTML form for creating a health facility order
func HealthFacilityOrderPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/health_facility_order.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}

// Serve the HTML form for creating a new vaccine
func AddVaccinePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/add_vaccine.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}

// Serve the HTML form for manufacturer's dashboard
func Manufacturerdashboard(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/manufacturer.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}

// func CreateDistributorOrder(w http.ResponseWriter, r *http.Request) {
// 	var transaction blockchain.VaccineTransaction
// 	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Printf("could not create distributor order: %v", err)
// 		w.Write([]byte("could not create distributor order"))
// 		return
// 	}

// 	if transaction.TransactionType != "DistributorToManufacturer" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("invalid transaction type for distributor order"))
// 		return
// 	}

// 	blockchain.BlockChain.AddBlock(transaction)
// 	resp, err := json.MarshalIndent(transaction, "", " ")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Printf("could not marshal payload: %v", err)
// 		w.Write([]byte("could not save distributor order"))
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(resp)
// }

// func CreateHealthFacilityOrder(w http.ResponseWriter, r *http.Request) {
// 	var transaction blockchain.VaccineTransaction
// 	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Printf("could not create health facility order: %v", err)
// 		w.Write([]byte("could not create health facility order"))
// 		return
// 	}

// 	if transaction.TransactionType != "FacilityToDistributor" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("invalid transaction type for health facility order"))
// 		return
// 	}

// 	blockchain.BlockChain.AddBlock(transaction)
// 	resp, err := json.MarshalIndent(transaction, "", " ")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Printf("could not marshal payload: %v", err)
// 		w.Write([]byte("could not save health facility order"))
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(resp)
// }

func AddVaccineHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		// Generate the next vaccine ID
		vaccineIDCounter++
		newVaccineID := fmt.Sprintf("vaccine%d", vaccineIDCounter)
		quantity, _ := strconv.Atoi(r.FormValue("quantity"))
		vaccine := Vaccine{
			ID:              newVaccineID,
			Name:            r.FormValue("vaccine_name"),
			BatchNumber:     r.FormValue("batch_no"),
			ManufactureDate: r.FormValue("manufacture_date"),
			ExpiryDate:      r.FormValue("expiry_date"),
			Quantity:        quantity,
		}

		availableVaccines = append(availableVaccines, vaccine)

		// Create a new transaction using the form values
		transaction := blockchain.VaccineTransaction{
			OrderID:        newVaccineID,
			IsGenesis:      false, // Set this according to your logic
			Details:        fmt.Sprintf("Vaccine: %s, Batch: %s", r.FormValue("vaccine_name"), r.FormValue("batch_no")),
			Manufacturer:   r.FormValue("manufacturer"),
			Distributor:    r.FormValue("distributor"),
			HealthFacility: "",
			AdministeredTo: "",
			Status:         "Manufactured",
			BatchNo:        r.FormValue("batch_no"),
			Quantity:       quantity,
			Timestamp:      time.Now().String(),
		}
		// Add the new transaction to the blockchain
		blockchain.BlockChain.AddBlock(transaction)
		http.Redirect(w, r, "/manufacturer-dashboard", http.StatusSeeOther)
	} else {
		http.Error(w, "Bad Method Request", http.StatusBadRequest)
	}
}
