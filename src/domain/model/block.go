package model

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type Blocker interface {
	CalculateHash() string
	Mine(int)
}

type block struct {
	Data         interface{} `json:"data"`
	Hash         string      `json:"hash"`
	PreviousHash string      `json:"previousHash"`
	Timestamp    time.Time   `json:"timestamp"`
	ProofOfWork  uint32      `json:"proofOfWork"`
}

func NewBlock(previousHash string, data interface{}) *block {
	b := block{
		Data:         data,
		Hash:         "",
		PreviousHash: previousHash,
		Timestamp:    time.Now(),
		ProofOfWork:  0,
	}

	return &b
}

func (b *block) CalculateHash() string {
	hash1 := sha256.New()

	dataBytes, err := json.Marshal(b.Data)
	if err != nil {
		log.Fatalln("Block=>CalculateHash() - failure to marshal data", err)
	}

	hashString := b.PreviousHash + string(dataBytes) + b.Timestamp.GoString() + strconv.Itoa(int(b.ProofOfWork))

	hash1.Write([]byte(hashString))

	return fmt.Sprintf("%x", hash1.Sum(nil))
}

func (b *block) Mine(difficulty int) {
	zeros := strings.Repeat("0", difficulty)

	for !strings.HasPrefix(b.Hash, zeros) {
		b.Hash = b.CalculateHash()
		b.ProofOfWork += 1
	}
}
