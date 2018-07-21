package main

import (
	"math/big"
	"bytes"
	"fmt"
	"math"
	"crypto/sha256"
)

const targetBits = 10

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	return &ProofOfWork{b, target}
}

func (pow *ProofOfWork) PrepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce))},
		[]byte{})
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < math.MaxInt64 {
		data := pow.PrepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	println("\n")
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.PrepareData(pow.block.nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	//   -1 if x <  y
	//    0 if x == y
	//   +1 if x >  y
	//   func (x *Int) Cmp(y *Int) (r int) {
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}
