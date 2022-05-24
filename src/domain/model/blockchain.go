package model

type Blockchainer interface {
	AddBlock(data interface{})
	IsValid() bool
}

type Blockchain struct {
	Chain []*block
}

func NewBlockchain() *Blockchain {
	bc := Blockchain{}

	genesisBlock := NewBlock("0", struct {
		isGenesis bool
	}{
		isGenesis: true,
	})

	bc.Chain = append(bc.Chain, genesisBlock)

	return &bc
}

func (bc *Blockchain) AddBlock(data interface{}) {
	lastBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := NewBlock(lastBlock.PreviousHash, data)

	newBlock.Mine(2)

	bc.Chain = append(bc.Chain, newBlock)
}

func (bc *Blockchain) IsValid() bool {
	chainLength := len(bc.Chain)

	for i := 1; i < chainLength; i++ {
		curr := bc.Chain[i]
		prev := bc.Chain[i-1]

		isNotCurrMatch := curr.Hash == curr.CalculateHash()
		isNotPrevMatch := curr.PreviousHash == prev.Hash

		if isNotCurrMatch || isNotPrevMatch {
			return false
		}
	}

	return true
}
