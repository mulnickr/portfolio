package posts

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"rmulnick.dev/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Post struct {
	ID       string `json:"id,omitempty" firestore:"id"`
	Title    string `json:"title" firestore:"title"`
	Date     string `json:"date,omitempty" firestore:"date"`
	Intro    string `json:"intro,omitempty" firestore:"intro"`
	Body     string `json:"body" firestore:"body"`
	ImageURL string `json:"imageurl" firestore:"imageurl"`
}

const (
	INTRO_LENGTH int = 100
)

func NewPost(ctx *gin.Context) {
	var request *Post
	erro := ctx.BindJSON(&request)
	if erro != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to bind to request body"})
		return
	}

	client, _ := database.ConfigureFirestore()
	defer client.Close()

	post := &Post{Title: request.Title, Body: request.Body, ImageURL: request.ImageURL}
	post.ID = uuid.NewString()
	post.Date = time.Now().Format("2006-01-02") //YYYY-MM-DD
	post.Intro = fmt.Sprintf("%v...", generateIntroParagraph(post.Body, INTRO_LENGTH))

	_, _, err := client.Collection("posts").Add(ctx, post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add new post."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": post})
}

func GetPost(ctx *gin.Context) {
	postId := ctx.Query("id")
	if postId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad request."})
		return
	}

	post, err := getPostFromID(postId, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func GetAllPosts(ctx *gin.Context) {
	var posts []*Post

	client, _ := database.ConfigureFirestore()
	defer client.Close()

	iter := client.Collection("posts").Documents(ctx)
	defer iter.Stop()

	docs, err := iter.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving posts"})
	}

	for _, doc := range docs {
		var post *Post

		if err := doc.DataTo(&post); err != nil {
			//this should probably just log the errors and allow returning valid posts
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error mapping data"})
			return
		}

		posts = append(posts, post)
	}

	ctx.JSON(http.StatusOK, posts)
}

func UpdatePost(ctx *gin.Context) {
	var post *Post
	err := ctx.BindJSON(&post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error mapping request body"})
		return
	}

	post.Intro = fmt.Sprintf("%v...", generateIntroParagraph(post.Body, INTRO_LENGTH))

	client, _ := database.ConfigureFirestore()
	defer client.Close()

	doc, err := client.Collection("posts").Where("id", "==", &post.ID).Documents(ctx).Next()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "database connection error", "err": err.Error()})
		return
	}

	if !doc.Exists() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "post not found"})
		return
	}

	_, erro := client.Collection("posts").Doc(doc.Ref.ID).Set(ctx, post)
	if erro != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error updating records: %v", erro)})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func DeletePost(ctx *gin.Context) {
	postId := ctx.Query("id")
	if postId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad request."})
		return
	}

	client, _ := database.ConfigureFirestore()
	defer client.Close()

	iter := client.Collection("posts").Where("id", "==", postId).Documents(ctx)
	defer iter.Stop()

	doc, err := iter.Next()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving posts"})
		return
	}

	if !doc.Exists() {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot find specified document"})
		return
	}

	_, erro := doc.Ref.Delete(ctx)
	if erro != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting post"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "post deleted."})
}

func getPostFromID(id string, ctx *gin.Context) (*Post, error) {
	client, _ := database.ConfigureFirestore()
	defer client.Close()

	var post *Post
	iter := client.Collection("posts").Where("id", "==", id).Documents(ctx)
	defer iter.Stop()
	docs, err := iter.GetAll()

	if err != nil || len(docs) == 0 {
		return nil, err
	}

	if len(docs) > 1 {
		return nil, errors.New("database conflic")
	}

	if err := docs[0].DataTo(&post); err != nil {
		return nil, errors.New("data mapping error")
	}

	if post == nil {
		return nil, errors.New("error retrieving post")
	}

	return post, nil
}

func generateIntroParagraph(body string, length int) string {
	// handle this better?
	replacer := strings.NewReplacer("#", "", "__", "", "---", "", "**", "")
	return replacer.Replace(body[:length])
}
