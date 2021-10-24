package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running bthesisAssistant server...")

	r := gin.Default()
	r.GET("/ping/:id", handlePing)
	r.GET("/ping", handlePing)
	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func handlePing(c *gin.Context) {
	id := c.Param("id")
	query := c.Query("q")
	c.JSON(http.StatusOK, gin.H{
		"message": "понг",
		"id":      id,
		"query":   query,
	})

}
