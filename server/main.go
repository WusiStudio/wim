package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	c1 := make(chan int)
	go func() {
		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c1 <- 1
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		r.Run() // listen and serve on 0.0.0.0:8080
	}()

	go func() {
		r := gin.Default()
		r.GET("/ping2", func(c *gin.Context) {
			c1 <- 2
			c.JSON(200, gin.H{
				"message": "pong2",
			})
		})
		r.Run(":8090") // listen and serve on 0.0.0.0:8080
	}()

	for {
		i := <-c1
		log.Print("---->", i)
	}
}
