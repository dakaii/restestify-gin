#Middleware for GIN and RESTestify

Example usage
```go
func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.New()

	r.Use(gin.Logger())
	// Add RESTestify as a middleware.
	// Remeber to spesify your api key
	r.Use(restestify.Logger("api key"))

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
```

TODO's
------
- Bundle requests to cut down on outgoing requests.