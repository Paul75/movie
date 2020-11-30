package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := NewMokeDB()
	initRoutes(r, db)
	r.Run()
}

func initRoutes(r *gin.Engine, db *MokeDB) {
	r.GET("/movies", func(c *gin.Context) {
		c.JSON(200, db.Movies)
	})
}
