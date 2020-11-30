package main

import (
	"net/http"

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
	r.POST("/movies", sm.Post)
}

type ServiceMovie struct {
	db *MokeDB
}

func (sm *ServiceMovie) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, sm.db.Movies)
}

func (sm *ServiceMovie) Post(ctx *gin.Context) {
	var m Movie
	if err := ctx.BindJSON(&m); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	newMovie := NewMovie(m.Title, m.Description, m.CreationDate)
	sm.db.Movies[newMovie.ID] = newMovie
	ctx.JSON(http.StatusOK, newMovie)
}
