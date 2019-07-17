package main

import (
	"log"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain_%s.db"
const blocksBucket = "blocks"

// Blockchain Blockchain
type Blockchain struct {
	// 最新的区块
	tip []byte
	db  *bolt.DB
}

// AddBlock AddBlock
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})
	if err!=nil{
		log.Panic(err)
	}
	newBlock := NewBlock(data, lastHash)
	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			return err
		}
		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			return err
		}
		bc.tip = newBlock.Hash
		return nil
	})
}

// NewGenesisBlock 创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlockchain 新建区块链
func NewBlockchain() *Blockchain {
	var tip []byte // 最后一个区块
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				return err
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				return err
			}
			err = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
			if err != nil {
				return err
			}
		}
		tip = b.Get([]byte("l"))
		return nil
	})
	bc := Blockchain{tip, db}
	return &bc
}

// Iterator Iterator
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}
	return bci
}
