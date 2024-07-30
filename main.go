package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"alpha-beta/blockchain"
	"alpha-beta/handlers"

	"github.com/gorilla/mux"
)

// Main function to set up routes and start the server
func main() {
	blockchain.BlockChain = blockchain.NewBlockchain()

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Serve HTML templates
	r.HandleFunc("/", handlers.BlockchainPage).Methods("GET")
	r.HandleFunc("/create-distributor-order", handlers.DistributorOrderPage).Methods("GET")
	r.HandleFunc("/create-health-facility-order", handlers.HealthFacilityOrderPage).Methods("GET")
	r.HandleFunc("/create-new-vaccine", handlers.NewVaccinePage).Methods("GET")

	// r.HandleFunc("/", getBlockchain).Methods("GET")
	r.HandleFunc("/distributor-order", handlers.CreateDistributorOrder).Methods("POST")
	r.HandleFunc("/health-facility-order", handlers.CreateHealthFacilityOrder).Methods("POST")
	r.HandleFunc("/new-vaccine", handlers.NewVaccine).Methods("POST")

	go func() {
		for _, block := range blockchain.BlockChain.Blocks {
			fmt.Printf("Prev. hash: %x\n", block.PrevHash)
			bytes, _ := json.MarshalIndent(block.Data, "", " ")
			fmt.Printf("Data: %v\n", string(bytes))
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()
		}
	}()
	log.Println("Listening on port 3000")

	log.Fatal(http.ListenAndServe(":3000", r))
}
