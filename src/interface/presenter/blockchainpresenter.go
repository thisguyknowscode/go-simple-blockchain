package presenter

import (
	"github.com/gin-gonic/gin"
	"github.com/thisguyknowscode/go-simple-blockchain/src/domain/model"
	"log"
)

func BlockchainHander(ctx *gin.Context) {
	data := "Give me the money, Lebowski"
	bc := model.NewBlockchain()

	bc.AddBlock(data)

	if ok := bc.IsValid(); !ok {
		log.Fatalln("BlockChainHandler() - blockchain validity check failed")
	}

	ctx.JSON(200, bc)
}
