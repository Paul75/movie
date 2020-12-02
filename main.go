package main

import (
	"github.com/gin-gonic/gin"

	"movie/concurrency"
	"movie/db/moke"
	"movie/service"
)

func main() {
	concurrency.Service()

	r := gin.Default()
	db := moke.NewMokeDB()
	//db := sqlite.New()
	service.InitRoutes(r, db)
	r.Run()
}
