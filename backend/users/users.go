package users

import (
	"context"
	"errors"
	"fmt"

	//"log"
	"net/http"
	"strings"
	"time"

	"rmulnick.dev/auth"
	"rmulnick.dev/database"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	TOKEN_LENGTH = 24
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Profile struct {
	ArticlesRead map[string][]string `firestore:"articles_read,omitempty" json:"articles_read,omitempty"`
	PracticeTime map[string][]string `firestore:"practice_time,omitempty" json:"practice_time,omitempty"`
}

type User struct {
	Email    string `firestore:"email" json:"email"`
	Username string `firestore:"username" json:"username"`
	Password string `firestore:"password" json:"-"`

	CreatedAt string `firestore:"createdat" json:"createdat"`
}

type UserSession struct {
	Permissions string
	LoginTime   string
	Token       string
}

func NewUserSession(permissions, loginTime, token string) UserSession {
	return UserSession{Permissions: permissions, LoginTime: loginTime, Token: token}
}

// gin routes
func Login(ctx *gin.Context) {
	var requestBody *LoginRequest
	ctx.BindJSON(&requestBody)

	client, _ := database.ConfigureFirestore()
	defer client.Close()

	user, err := getUserByEmail(requestBody.Email, client, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error logging in"})
		return
	}

	auth.SetToken(ctx, user.Email)
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "token": auth.GetCookie(ctx, "Token")})
}

func Logout(ctx *gin.Context) {
	_, err := ctx.Cookie("Token")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	auth.ClearToken(ctx)
	auth.DeleteCookie(ctx, "Token") // cleanup cookies temporarily...
	auth.DeleteCookie(ctx, "token")
	ctx.JSON(http.StatusOK, gin.H{"success": "logout successful"})
}

func Refresh(ctx *gin.Context) {
	auth.RefreshToken(ctx)
	ctx.JSON(http.StatusOK, gin.H{"success": "auth cookie refreshed"})
}

func Register(ctx *gin.Context) {
	var responseBody *LoginRequest
	ctx.BindJSON(&responseBody)

	client, _ := database.ConfigureFirestore()
	defer client.Close()

	password, err := bcrypt.GenerateFromPassword([]byte(responseBody.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Encryption error."})
		return
	}

	user, err := createUser(responseBody.Email, string(password), client, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to register: %v", err)})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": fmt.Sprintf("Account created: %v", user.Email)})
}

// database functions
func checkUserExists(email string, client *firestore.Client, ctx context.Context) bool {
	user, err := getUserByEmail(email, client, ctx)
	return user != nil && err == nil
}

func createUser(email, password string, client *firestore.Client, ctx context.Context) (*User, error) {
	if checkUserExists(email, client, ctx) {
		return nil, errors.New("user already exists")
	}

	user := &User{Email: email, Username: strings.Split(email, "@")[0], Password: password}
	user.CreatedAt = time.Now().Format("2006-01-02") //YYYY-MM-DD

	_, _, erro := client.Collection("users").Add(ctx, user)
	if erro != nil {
		return nil, erro
	}

	return user, nil
}

func getUserByEmail(email string, client *firestore.Client, ctx context.Context) (*User, error) {
	var user *User
	iter := client.Collection("users").Where("email", "==", email).Documents(ctx)
	docs, err := iter.GetAll()
	if err != nil || len(docs) == 0 {
		return nil, errors.New("no user doc found")
	}

	if len(docs) > 1 {
		return nil, errors.New("database error: too many users with the same identifier")
	}

	if err := docs[0].DataTo(&user); err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("unresolved 'get user' error")
	}

	return user, nil
}
