package main

import (
	"github.com/gin-gonic/gin"

	"movie/concurrency"
	"movie/db/moke"
	"movie/middleware"
	"movie/service"
)

func main() {
	concurrency.Service()
	middleware.MySigningKey = []byte("als;dkjf;lasdjfl;saj")
	r := gin.Default()
	db := moke.NewMokeDB()
	//db := sqlite.New()
	service.InitRoutes(r, db)
	r.Run()
}
