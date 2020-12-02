package service

import (
	"movie/db"
	"movie/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, db db.DB) {
	sm := ServiceMovie{db}
	moviesPath := r.Group("/movies")
	moviesPath.Use(middleware.JWT("AllYourBase"))
	moviesPath.GET("", sm.Get)
	moviesPath.POST("", sm.Post)
	moviesPath.GET("/:uuid", sm.GetByUUID)
	moviesPath.DELETE("/:uuid", sm.Delete)
	moviesPath.PATCH("/:uuid", sm.Update)
	// service User
	su := ServiceUser{db}
	r.POST("/users", su.Post)
	r.GET("/users", su.Get)
	r.GET("/users/:uuid", su.GetByUUID)
	r.DELETE("/users/:uuid", su.Delete)
	r.PATCH("/users/:uuid", su.Update)
	r.POST("/login", su.Login)
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
