package service

import (
	"movie/db"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, db db.DB) {
	sm := ServiceMovie{db}
	r.GET("/movies", sm.Get)
	r.POST("/movies", sm.Post)
	r.GET("/movies/:uuid", sm.GetByUUID)
	r.DELETE("/movies/:uuid", sm.Delete)
	r.PATCH("/movies/:uuid", sm.Update)
	// service User
	su := ServiceUser{db}
	r.POST("/users", su.Post)
	r.GET("/users", su.Get)
	r.GET("/users/:uuid", su.GetByUUID)
	r.DELETE("/users/:uuid", su.Delete)
	r.PATCH("/users/:uuid", su.Update)
	// service Actor
	sa := ServiceActor{db}
	r.POST("/actors", sa.Post)
	r.GET("/actors", sa.Get)
	r.GET("/actors/:uuid", sa.GetByUUID)
	r.DELETE("/actors/:uuid", sa.Delete)
	r.PATCH("/actors/:uuid", sa.Update)
	// service Media
	smed := ServiceMedia{db}
	r.POST("/medias", smed.Post)
	r.GET("/medias", smed.Get)
	r.GET("/medias/:uuid", smed.GetByUUID)
	r.DELETE("/medias/:uuid", smed.Delete)
	r.PATCH("/medias/:uuid", smed.Update)
}
