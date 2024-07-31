package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"
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
	OrderID        string `json:"order_id"`
	IsGenesis      bool   `json:"is_genesis"`
	Details        string `json:"details"`
	Manufacturer   string `json:"manufacturer"`
	Distributor    string `json:"distributor"`
	HealthFacility string `json:"health_facility"`
	AdministeredTo string `json:"administered_to"`
	Status         string `json:"status"`
	BatchNo        string `json:"batch_no"`
	Quantity       int    `json:"quantity"`
	Timestamp      string `json:"timestamp"`
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

// Blockchain represents a series of Blocks
type Blockchain struct {
	Blocks []*Block
}

var BlockChain *Blockchain

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data VaccineTransaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	block := CreateBlock(prevBlock, data)

	if validBlock(block, prevBlock) {
		bc.Blocks = append(bc.Blocks, block)
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
	return b.Hash == hash
}

// GetBlockchain returns the current blockchain as JSON
func GetBlockchain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the correct content type
	jbytes, err := json.MarshalIndent(BlockChain.Blocks, "", "    ")
	if err != nil {
		http.Error(w, "Failed to marshal blockchain data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jbytes)
}
