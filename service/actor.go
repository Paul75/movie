package service

import (
	"movie/db"
	"movie/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceActor struct {
	db db.DB
}

func (su *ServiceActor) Get(ctx *gin.Context) {
	us, err := su.db.GetUsers()
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, us)
}

func (su *ServiceActor) GetByUUID(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	u, err := su.db.GetUserByUUID(uuid)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, u)
}

func (su *ServiceActor) Post(ctx *gin.Context) {
	var u model.User
	if err := ctx.BindJSON(&u); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user := model.NewUser(u.FirstName, u.LastName, u.Email, u.Password)
	_, err := su.db.AddUser(user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (su *ServiceActor) Delete(ctx *gin.Context) {
	err := su.db.DeleteUser(ctx.Param("uuid"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (su *ServiceActor) Update(ctx *gin.Context) {
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
