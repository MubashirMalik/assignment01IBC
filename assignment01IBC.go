package assignment01IBC

import (
	"crypto/sha256"
	"fmt"
)

type BlockData struct {
	Transactions []string
}

type Block struct {
	Data        BlockData
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

func SearchString(Transactions []string, oldTrans string) int {
	for i := 0; i < len(Transactions); i++ {
		if Transactions[i] == oldTrans {
			return i
		}
	}
	return -1
}

func CalculateHash(inputBlock *Block) string {
	transactions := inputBlock.Data.Transactions[0]
	for i := 1; i < len(inputBlock.Data.Transactions); i++ {
		transactions += inputBlock.Data.Transactions[i]
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(transactions)))
}

func InsertBlock(dataToInsert BlockData, chainHead *Block) *Block {
	var newBlock *Block = new(Block)
	newBlock.Data = dataToInsert

	if chainHead == nil {
		newBlock.PrevPointer = nil
		newBlock.PrevHash = "g3n3515"
	} else {
		newBlock.PrevPointer = chainHead
		newBlock.PrevHash = chainHead.CurrentHash
	}

	newBlock.CurrentHash = CalculateHash(newBlock)
	return newBlock
}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	for block := chainHead; block != nil; block = block.PrevPointer {
		index := SearchString(block.Data.Transactions, oldTrans)
		if index != -1 {
			block.Data.Transactions[index] = newTrans
			break
		}
	}
}

func ListBlocks(chainHead *Block) {
	for block := chainHead; block != nil; block = block.PrevPointer {
		fmt.Println("\nBlock - ")
		fmt.Println("Data:        ", block.Data)
		fmt.Println("PrevPointer: ", &block.PrevPointer)
		fmt.Println("PrevHash:    ", block.PrevHash)
		fmt.Println("CurrentHash: ", block.CurrentHash)
	}
}

func VerifyChain(chainHead *Block) {
	for block := chainHead; block != nil; block = block.PrevPointer {
		if block.CurrentHash != CalculateHash(block) {
			fmt.Println("\nInvalid Chain! Current Hash of block is tampered.")
			break
		}
	}
}
