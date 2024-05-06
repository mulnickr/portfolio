package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	MAX_AGE = 24 * time.Hour // Cookie lasts for one day
	DELETE  = -1             // set cookie age to a time in the past
	DOMAIN  = "localhost"    // production domain: rmulnick.dev
)

var (
	secretKey = []byte("secret-key") // initialize with ENV variable
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func ValidateToken(ctx *gin.Context) bool {
	token, err := ctx.Cookie("Token")
	if err != nil {
		fmt.Printf("Error fetching token: %v\n", err)
		return false
	}

	claims := &Claims{}
	parse, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) { return secretKey, nil })
	if err != nil || !parse.Valid {
		fmt.Printf("Error parsing token, or invalid token: %v\n", err)
		return false
	}

	return true
}

func SetToken(ctx *gin.Context, email string) {
	expiration := time.Now().Add(MAX_AGE)
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"": ""})
		return
	}

	/* 	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    "Token",
		Value:   tokenString,
		Expires: expiration,
	}) */
	SetCookie(ctx, "Token", tokenString, expiration)
}

func ClearToken(ctx *gin.Context) {
	DeleteCookie(ctx, "Token")
}

func RefreshToken(ctx *gin.Context) {
	cookie, err := ctx.Cookie("Token")
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(cookie, claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	expiration := time.Now().Add(MAX_AGE)
	claims.ExpiresAt = jwt.NewNumericDate(expiration)

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString(secretKey)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	SetCookie(ctx, "Token", tokenString, expiration)

	/* 	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    "Token",
		Value:   tokenString,
		Expires: expiration,
	}) */
}

// alternate cookie configuration with Gin
func SetCookie(ctx *gin.Context, name, value string, expiration time.Time) {
	ctx.SetCookie(name, value, int(expiration.Unix()), "/", DOMAIN, true, true)
}

func GetCookie(ctx *gin.Context, name string) string {
	cookie, err := ctx.Cookie(name)
	if err != nil {
		return ""
	}

	return cookie
}

func RefreshCookie(ctx *gin.Context, name string) {
	ctx.SetCookie(name, GetCookie(ctx, name), int(MAX_AGE), "/", DOMAIN, true, true)
}

func DeleteCookie(ctx *gin.Context, name string) {
	ctx.SetCookie(name, "", DELETE, "/", DOMAIN, true, true)
}
