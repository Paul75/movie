package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"movie/concurrency"
	"movie/db/moke"
	"movie/middleware"
	"movie/service"
)

type Config struct {
	MySigningKey string
}

var conf Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.SetDefault("MySigningKey", "AFnakljdfhskjhdfsffk")
	conf.MySigningKey = viper.GetString("MySigningKey")
}

func main() {

	concurrency.Service()
	middleware.MySigningKey = []byte(conf.MySigningKey)
	r := gin.Default()
	db := moke.NewMokeDB()
	//db := sqlite.New()
	service.InitRoutes(r, db)
	r.Run()
}
