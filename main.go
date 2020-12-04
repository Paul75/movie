package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"movie/cache/redis"
	"movie/db"
	"movie/db/moke"
	"movie/db/postgres"
	"movie/db/sqlite"
	"movie/middleware"
	"movie/service"
)

type Config struct {
	Env          string
	MySigningKey string
	PostgresDNS  string
	Redis        struct {
		DNS string
		Exp int
	}
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
	conf.Env = viper.GetString("env")
	viper.SetDefault("MySigningKey", "AFnakljdfhskjhdfsffk")
	conf.MySigningKey = viper.GetString("MySigningKey")
	conf.Redis.DNS = viper.GetString("redis.dns")
	conf.Redis.Exp = viper.GetInt("redis.exp")
	conf.PostgresDNS = viper.GetString("postgresDNS")
}

func main() {

	// concurrency.Service()
	middleware.MySigningKey = []byte(conf.MySigningKey)
	r := gin.Default()
	var db db.DB
	switch conf.Env {
	case "local":
		db = sqlite.New()
	case "pord":
		db = postgres.New(conf.PostgresDNS)
	default:
		db = moke.NewMokeDB()

	}
	//db := sqlite.New()
	redisDB := redis.NewRedisDB(conf.Redis.DNS, conf.Redis.Exp)
	service.InitRoutes(r, db, redisDB)
	r.Run()
}
