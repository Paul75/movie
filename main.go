package main

import (
	"github.com/gin-gonic/gin"

	"movie/db/moke"
	"movie/service"
)

func main() {
	r := gin.Default()
	db := moke.NewMokeDB()
	service.InitRoutes(r, db)
	r.Run()
}
