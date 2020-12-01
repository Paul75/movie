package service

import (
	"movie/db"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, db db.DB) {
	sm := ServiceMovie{db}
	r.GET("/movies", sm.Get)
	r.POST("/movies", sm.Post)
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
}
