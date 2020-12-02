package service

import (
	"movie/db"
	"movie/middleware"
	"movie/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceUser struct {
	db db.DB
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
	if err := ctx.BindJSON(&data); err != nil {
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

type LoginPayload struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

func (su *ServiceUser) Login(ctx *gin.Context) {
	var payload LoginPayload
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	u, err := su.db.GetUserByEmail(payload.Login)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if u.Password != model.HashPass(payload.Pass) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	jwtValue, err := middleware.GetJWT(u.ID, u.Email)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"jwt": jwtValue})
}
