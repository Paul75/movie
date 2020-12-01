package service

import (
	"movie/db"
	"movie/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceMovie struct {
	db db.DB
}

func (sm *ServiceMovie) Get(ctx *gin.Context) {
	ms, err := sm.db.GetMovies()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, ms)
}

func (sm *ServiceMovie) Post(ctx *gin.Context) {
	var m model.Movie
	if err := ctx.BindJSON(&m); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	newMovie := model.NewMovie(m.Title, m.Description, m.CreationDate)
	_, err := sm.db.AddMovie(newMovie)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, newMovie)
}
