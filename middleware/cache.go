package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"movie/cache"
)

func Cache(cache cache.CacheDB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := cache.Get(ctx.Request.URL.String())
		if err != nil || len(data) == 0 {
			fmt.Println("create cache")
			ctx.Next()
		} else {
			fmt.Println("try to wite cache in the response")
			ctx.Header("Content-Type", "application/json")
			ctx.Writer.WriteHeader(http.StatusOK)
			ctx.Writer.WriteString(data)
			ctx.Abort()
		}
	}
}
