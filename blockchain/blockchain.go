package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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
	OrderID   string `json:"order_id"`
	IsGenesis bool   `json:"is_genesis"`
	Details   string `json:"details"`
}

// generateHash computes the hash for the block
func (b *Block) generateHash() {
	bytes, _ := json.Marshal(b.Data)
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
	bc.Blocks = append(bc.Blocks, block)
}

// GenesisBlock creates the first block in the blockchain
func GenesisBlock() *Block {
	return CreateBlock(&Block{}, VaccineTransaction{IsGenesis: true})
}

// NewBlockchain initializes a new blockchain with the genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}
