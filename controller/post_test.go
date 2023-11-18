package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, CreatePostHandler)
	body := `{
		"community_id": 1,
		"title": "test",
		"content": "just a test"
	}`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatal("json.Unmarshal w.Body Failed", err)
	}
	assert.Equal(t, res.Code, CodeNeedLogin)
}
