package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	"rmulnick.dev/auth"
	"rmulnick.dev/posts"
	"rmulnick.dev/users"
)

func main() {
	router := gin.Default()

	// CORS policy config
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://rmulnick.dev", "http://localhost:4200", "https://*.rmulnick.dev"},
		AllowMethods:     []string{"GET", "PUT", "POST"},
		AllowHeaders:     []string{"Token", "Content-Type", "Origin", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	}))

	router.GET("/health", health)

	userGroup := router.Group("/users")
	userGroup.POST("/login", users.Login)
	userGroup.GET("/logout", users.Logout)
	userGroup.POST("/register", users.Register)
	userGroup.GET("/refresh", users.Refresh)

	postGroup := router.Group("/posts")
	postGroup.GET("/", posts.GetAllPosts)
	postGroup.GET("/p", posts.GetPost)

	// requires authentication
	editGroup := postGroup.Group("/edit")
	editGroup.Use(Authenticator())

	editGroup.POST("/new", posts.NewPost)
	editGroup.PUT("/update", posts.UpdatePost)
	editGroup.GET("/delete", posts.DeletePost)

	router.Run(":5000")
}

func health(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{"health": "active"})
}

func Authenticator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if valid := auth.ValidateToken(ctx); !valid {
			cookie := auth.GetCookie(ctx, "Token")
			fmt.Printf("Unauthorized: %v\n", cookie)
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "unauthorized request"})
			return
		}

		ctx.Next()
	}
}
