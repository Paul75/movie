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
	// service User
	su := ServiceUser{db}
	r.POST("/users", su.Post)
	r.GET("/users", su.Get)
	r.GET("/users/:uuid", su.GetByUUID)
	r.DELETE("/users/:uuid", su.Delete)
	r.PATCH("/users/:uuid", su.Update)
}

type ServiceUser struct {
	db *MokeDB
}

func (su *ServiceUser) Get(ctx *gin.Context) {
	us, err := su.db.GetUsers()
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, us)
}

func (su *ServiceUser) GetByUUID(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	u, err := su.db.GetUserByUUID(uuid)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, u)
}

func (su *ServiceUser) Post(ctx *gin.Context) {
	var u User
	if err := ctx.BindJSON(&u); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user := NewUser(u.FirstName, u.LastName, u.Email, u.Password)
	_, err := su.db.AddUser(user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (su *ServiceUser) Delete(ctx *gin.Context) {
	err := su.db.DeleteUser(ctx.Param("uuid"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (su *ServiceUser) Update(ctx *gin.Context) {
	data := make(map[string]interface{})
	if err := ctx.BindJSON(data); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	u, err := su.db.UpdateUser(ctx.Param("uuid"), data)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, u)
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
