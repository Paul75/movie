package service

import (
	"movie/db"
	"movie/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceMedia struct {
	db db.DB
}

func (su *ServiceMedia) Get(ctx *gin.Context) {
	us, err := su.db.GetMedias()
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, us)
}

func (su *ServiceMedia) GetByUUID(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	u, err := su.db.GetMediaByUUID(uuid)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, u)
}

func (su *ServiceMedia) Post(ctx *gin.Context) {
	var m model.Media
	if err := ctx.BindJSON(&m); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	Media := model.NewMedia(m.Title, m.URI, m.Kind)
	_, err := su.db.AddMedia(Media)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, Media)
}

func (su *ServiceMedia) Delete(ctx *gin.Context) {
	err := su.db.DeleteMedia(ctx.Param("uuid"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (su *ServiceMedia) Update(ctx *gin.Context) {
	data := make(map[string]interface{})
	if err := ctx.BindJSON(&data); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	u, err := su.db.UpdateMedia(ctx.Param("uuid"), data)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, u)
}
