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
	sm := ServiceMovie{db}
	r.GET("/movies", sm.Get)
}

type ServiceMovie struct {
	db *MokeDB
}

func (sm *ServiceMovie) Get(ctx *gin.Context) {
	ctx.JSON(200, sm.db.Movies)
}
