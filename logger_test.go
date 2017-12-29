package restestify

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func TestLogger(t *testing.T) {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(Logger("demo"))

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
