package main

import (
	"github.com/boltdb/bolt"
	"os"
)

const dbfile = "blockChainDb.db"
const blockBucket  = "block"
const lasthash  = "lastHash"

//改造区块结构，使用数组来储存所有的区块
type BlockChain struct {
	//blocks []*Block
	db *bolt.DB
	lastHash []byte
}

//创建区块链实例，并且添加第一个创世快
func NewBlockChain() *BlockChain {
	//return &BlockChain{[]*Block{NewGenesisBlock()}}

	//func Open(path string, mode os.FileMode, options *Options) (*DB, error) {
	db, err := bolt.Open(dbfile, 0600, nil)
	CheckErr(err)
	var lastHash []byte

	//func (db *DB) Update(fn func(*Tx) error) error {
	db.Update(func(tx *bolt.Tx) error {
		//func (tx *Tx) Bucket(name []byte) *Bucket {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket != nil{
			//读取lastHash
			//func (b *Bucket) Get(key []byte) []byte {
			lastHash = bucket.Get([]byte(lasthash))
		}else {
			//1.创建bucket
			//2.写数据
			genesis := NewGenesisBlock()
			//func (tx *Tx) CreateBucket(name []byte) (*Bucket, error) {
			bucket, err := tx.CreateBucket([]byte(blockBucket))
			CheckErr(err)

			//func (b *Bucket) Put(key []byte, value []byte) error {
			err = bucket.Put(genesis.Hash, genesis.Serialize())
			CheckErr(err)

			err = bucket.Put([]byte(lasthash), genesis.Hash)
			CheckErr(err)
			lastHash = genesis.Hash
		}
		return nil
	})
	return &BlockChain{db, lastHash}
}

//添加区块操作
func (bc *BlockChain) AddBlock(data string) {
	var prevBlockHash []byte
	err := bc.db.View(func(tx *bolt.Tx) error {
		//取出最后一个区块，目的是得到其hash值
		bucket := tx.Bucket([]byte(blockBucket))
		lastHash := bucket.Get([]byte(lasthash))
		prevBlockHash = lastHash
		return nil
	})
	CheckErr(err)
	block := NewBlock(data, prevBlockHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		//改造一个新的将要添加到区块链的区块
		err := bucket.Put(block.Hash, block.Serialize())
		CheckErr(err)

		err = bucket.Put([]byte(lasthash), block.Hash)
		CheckErr(err)
		bc.lastHash = block.Hash
		return nil
	})
	CheckErr(err)
}

type BlockChainIterator struct {
	db *bolt.DB
	currentHash []byte
}

func (bc *BlockChain)Iterator() *BlockChainIterator {
	return &BlockChainIterator{bc.db, bc.lastHash}
}

func (it *BlockChainIterator)Next() *Block{
	var block *Block
	err := it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil{
			os.Exit(1)
		}
		blockTmp := bucket.Get(it.currentHash)
		block = Deserialize(blockTmp)
		it.currentHash = block.PrevBlockHash
		return nil
	})
	CheckErr(err)
	return block
}


