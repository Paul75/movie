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

func (su *ServiceMovie) GetByUUID(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	m, err := su.db.GetMovieByUUID(uuid)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, m)
}

func (su *ServiceMovie) Delete(ctx *gin.Context) {
	err := su.db.DeleteMovie(ctx.Param("uuid"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (su *ServiceMovie) Update(ctx *gin.Context) {
	data := make(map[string]interface{})
	if err := ctx.BindJSON(&data); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	u, err := su.db.UpdateMovie(ctx.Param("uuid"), data)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, u)
}
