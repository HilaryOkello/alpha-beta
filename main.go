package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Block represents a single block in the blockchain
type Block struct {
	Pos       int
	Data      VaccineTransaction
	Timestamp string
	Hash      string
	PrevHash  string
}

// VaccineTransaction represents a transaction related to vaccine distribution
type VaccineTransaction struct {
	OrderID      string `json:"order_id"`
	ManufacturerID string `json:"manufacturer_id"`
	HealthFacilityID string `json:"health_facility_id"`
	VaccineDetails string `json:"vaccine_details"`
	TransactionType   string `json:"transaction_type"`
	IsGenesis    bool   `json:"is_genesis"`
}

// Vaccine represents a vaccine with details
type Vaccine struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	ExpiryDate  string `json:"expiry_date"`
	BatchNumber string `json:"batch_number"`
}

// generateHash computes the hash for the block
func (b *Block) generateHash() {
	// Get string representation of the Data
	bytes, _ := json.Marshal(b.Data)
	// Concatenate the dataset
	data := string(b.Pos) + b.Timestamp + string(bytes) + b.PrevHash
	hash := sha256.New()
	hash.Write([]byte(data))
	b.Hash = hex.EncodeToString(hash.Sum(nil))
}

// CreateBlock creates a new block with the provided data
func CreateBlock(prevBlock *Block, transaction VaccineTransaction) *Block {
	block := &Block{}
	block.Pos = prevBlock.Pos + 1
	block.Timestamp = time.Now().String()
	block.Data = transaction
	block.PrevHash = prevBlock.Hash
	block.generateHash()

	return block
}

// Blockchain represents a series of blocks
type Blockchain struct {
	blocks []*Block
}

var BlockChain *Blockchain

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data VaccineTransaction) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	block := CreateBlock(prevBlock, data)

	if validBlock(block, prevBlock) {
		bc.blocks = append(bc.blocks, block)
	}
}

// GenesisBlock creates the first block in the blockchain
func GenesisBlock() *Block {
	return CreateBlock(&Block{}, VaccineTransaction{IsGenesis: true})
}

// NewBlockchain initializes a new blockchain with the genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

// validBlock checks if the block is valid
func validBlock(block, prevBlock *Block) bool {
	if prevBlock.Hash != block.PrevHash {
		return false
	}
	if !block.validateHash(block.Hash) {
		return false
	}
	if prevBlock.Pos+1 != block.Pos {
		return false
	}
	return true
}

// validateHash verifies the hash of the block
func (b *Block) validateHash(hash string) bool {
	b.generateHash()
	if b.Hash != hash {
		return false
	}
	return true
}

// getBlockchain returns the current blockchain as JSON
func getBlockchain(w http.ResponseWriter, r *http.Request) {
	jbytes, err := json.MarshalIndent(BlockChain.blocks, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	io.WriteString(w, string(jbytes))
}

// writeBlock handles adding a new block to the blockchain
func writeBlock(w http.ResponseWriter, r *http.Request) {
	var transaction VaccineTransaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not write Block: %v", err)
		w.Write([]byte("could not write block"))
		return
	}

	// Ensure transaction type is specified
	if transaction.TransactionType != "DistributorToManufacturer" && transaction.TransactionType != "FacilityToDistributor" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid transaction type"))
		return
	}
	
	BlockChain.AddBlock(transaction)
	resp, err := json.MarshalIndent(transaction, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not marshal payload: %v", err)
		w.Write([]byte("could not write block"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// newVaccine handles creating a new vaccine record
func newVaccine(w http.ResponseWriter, r *http.Request) {
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

// Main function to set up routes and start the server
func main() {
	BlockChain = NewBlockchain()

	r := mux.NewRouter()
	r.HandleFunc("/", getBlockchain).Methods("GET")
	r.HandleFunc("/", writeBlock).Methods("POST")
	r.HandleFunc("/new-vaccine", newVaccine).Methods("POST")

	go func() {
		for _, block := range BlockChain.blocks {
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
