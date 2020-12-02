package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var MySigningKey []byte

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		val := ctx.Request.Header.Get("Authorization")
		if len(val) == 0 || !strings.Contains(val, "Bearer ") {
			log.Println("no vals or no Bearer found")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		vals := strings.Split(val, " ")
		if len(vals) != 2 {
			log.Println("result split not valid")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(vals[1], validateJWT)

		if err != nil {
			log.Println("error parsing JWT", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(MyCustomClaims); ok && token.Valid {
			fmt.Println(claims.ID, claims.Email)
		}
	}
}

func validateJWT(token *jwt.Token) (interface{}, error) {
	log.Println("try to pars the JWT")
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		log.Println("error parsing JWT")
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return MySigningKey, nil
}

type MyCustomClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GetJWT(id, email string) (string, error) {
	// Create the Claims
	claims := MyCustomClaims{
		id,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "movie API",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtValue, err := token.SignedString(MySigningKey)
	if err != nil {
		return "", err
	}
	return jwtValue, nil
}
