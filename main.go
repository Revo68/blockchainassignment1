package main

import (
	"blockchainassignment1/assignment01bca" // Correct import path
	"fmt"
)

func main() {
	// Creating new blocks
	block1 := assignment01bca.NewBlock("bob to alice", 1, "")
	block2 := assignment01bca.NewBlock("alice to bob", 2, block1.Hash)

	// Listing blocks
	fmt.Println("Listing all blocks in the blockchain:")
	assignment01bca.ListBlocks()

	// Use block2 (for example, print its details)
	fmt.Printf("Block 2 transaction: %s, Hash: %s\n", block2.Transaction, block2.Hash)

	// Verify the chain
	fmt.Println("Verifying blockchain integrity:")
	assignment01bca.VerifyChain()

	// Change a block and verify again
	fmt.Println("Tampering with the first block:")
	assignment01bca.ChangeBlock()

	fmt.Println("Verifying blockchain integrity again after tampering:")
	assignment01bca.VerifyChain()
}
