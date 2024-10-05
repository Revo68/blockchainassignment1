package assignment01bca

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

// Block structure
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

// Blockchain as a slice of blocks
var Blockchain []Block

// Function to calculate the hash of a block
func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	return fmt.Sprintf("%x", hash)
}

// Create a new block
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.Hash = CalculateHash(block.Transaction + strconv.Itoa(block.Nonce) + block.PreviousHash)
	Blockchain = append(Blockchain, *block)
	return block
}

// List all blocks in the blockchain
func ListBlocks() {
	for i, block := range Blockchain {
		fmt.Printf("Block #%d\n", i)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println("---------------")
	}
}

// Function to change a block (for demonstration)
func ChangeBlock() {
	if len(Blockchain) > 0 {
		Blockchain[0].Transaction = "tampered transaction"
		fmt.Println("First block transaction changed")
	}
}

// Verify the integrity of the blockchain
func VerifyChain() {
	for i := range Blockchain {
		if i == 0 {
			continue
		}
		if Blockchain[i].PreviousHash != Blockchain[i-1].Hash {
			fmt.Println("Blockchain integrity compromised at block", i)
			return
		}
	}
	fmt.Println("Blockchain is valid")
}
