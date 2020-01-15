package main

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

// MyBenchLogger
func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("benchmark test 1")
		status := c.Writer.Status()
		fmt.Println(status)
	}
}

// AuthRequired test
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("auth test 1")
		status := c.Writer.Status()
		fmt.Println(status)
		c.JSON(401, gin.H{
			"code": 401,
			"data": "未认证",
		})
		c.Abort()
	}
}

func main() {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Per route middleware, you can add as many as you desire.
	r.GET("/benchmark", MyBenchLogger(), func(c *gin.Context) {
		fmt.Println("benchmark func 1")
		c.JSON(200, gin.H{
			"code": http.StatusOK,
			"data": "benchmark func 1",
		})
	})

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(AuthRequired())
	{
		authorized.GET("/login", func(c *gin.Context) {
			fmt.Println("benchmark func login")
			c.JSON(200, gin.H{
				"code": http.StatusOK,
				"data": "benchmark func login",
			})
		})
		authorized.GET("/submit", func(c *gin.Context) {
			fmt.Println("benchmark func submit")
			c.JSON(200, gin.H{
				"code": http.StatusOK,
				"data": "benchmark func submit",
			})
		})
		authorized.GET("/read", func(c *gin.Context) {
			fmt.Println("benchmark func read")
			c.JSON(200, gin.H{
				"code": http.StatusOK,
				"data": "benchmark func read",
			})
		})

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", func(c *gin.Context) {
			fmt.Println("benchmark test 1")
			c.JSON(200, gin.H{
				"code": http.StatusOK,
				"data": "benchmark func testing",
			})
		})
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run("0.0.0.0:8080")
}

