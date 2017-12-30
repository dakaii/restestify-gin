package restestify

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"encoding/base64"
	"runtime"
	"sync"
	"os"
)

var (
	once     sync.Once
	apiKey   string
	hostname string
)

func Logger(key string) gin.HandlerFunc {
	// Init
	once.Do(func() {
		apiKey = key
		// Start workers
		JobQueue = make(chan Job)
		NewDispatcher(runtime.NumCPU()).Run()
		// Get the hostname
		hostname, _ = os.Hostname()
	})

	return func(c *gin.Context) {
		if key == "" {
			log.Printf("ERROR: invalid restestify middleware api key. recived: %s", key)
			// Continue request
			c.Next()
		} else {
			// Start timer
			start := time.Now()
			query := []byte(c.Request.URL.RawQuery)

			// Process request
			c.Next()

			// Stop timer
			end := time.Now()
			latency := end.Sub(start)

			clientIP := c.ClientIP()
			method := c.Request.Method
			statusCode := c.Writer.Status()

			// Push to worker
			JobQueue <- Job{Request: Request{
				Path:     c.Request.URL.Path,
				// Base64 encode the query string to escape any potential shady chars
				Query:    base64.StdEncoding.EncodeToString(query),
				Status:   statusCode,
				Latency:  latency.Nanoseconds(),
				ClientIp: clientIP,
				Method:   method,
				Time:     end.UTC(),
				Hostname: hostname,
			}}
		}
	}
}
