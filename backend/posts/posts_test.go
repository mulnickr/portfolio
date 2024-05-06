package posts

/*
import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewPost(t *testing.T) {
	t.Run("Creates a new post.", func(t *testing.T) {
		jsonBody := []byte(`{"id": "testID", "title": "test_title"}`)
		bodyReader := bytes.NewReader(jsonBody)

		request, _ := http.NewRequest(http.MethodPost, "/posts/new", bodyReader)
		response := httptest.NewRecorder()

		context, _ := gin.CreateTestContext(response)
		NewPost(context)

		got := response.Body.String()
		want := jsonBody
	})
}
*/
