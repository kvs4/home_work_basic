package main

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	v1 := server.Group("/v1")

	{
		v1.GET("/", func(c *gin.Context) {
			fmt.Println("GET body:", c.Request.Body)
			c.JSONP(200, "this is method GET")
		})

		v1.POST("/", func(c *gin.Context) {
			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				fmt.Println(fmt.Errorf("POST error read body: %w", err))
				return
			}
			fmt.Println("POST body:", string(body))
			c.JSONP(200, "this is method POST")
		})
	}

	server.Run(":8080")
}

/*func main() {
	fmt.Println("hi")
}*/
