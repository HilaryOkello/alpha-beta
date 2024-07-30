package main

import (
	"log"
	"net/http"

	"alpha-beta/blockchain"
	"alpha-beta/database"
	"alpha-beta/handlers"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()

	blockchain.BlockChain = blockchain.NewBlockchain()

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.HandleFunc("/create-health-facility-order", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/pharmacy.html")
	}).Methods("GET")

	r.HandleFunc("/inventory", handlers.InventoryPage).Methods("GET")

	r.HandleFunc("/health-facility-order", handlers.CreateHealthFacilityOrder).Methods("POST")

	log.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
