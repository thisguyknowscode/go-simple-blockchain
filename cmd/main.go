package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thisguyknowscode/go-simple-blockchain/src/interface/presenter"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/", presenter.BlockchainHander)

	err := r.Run()
	if err != nil {
		log.Fatalln("Application failed to run", err)
	}
}
