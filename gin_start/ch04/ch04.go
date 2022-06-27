package main

import "github.com/gin-gonic/gin"

type Person struct {
	ID   int    `uri:"id" binding:"required"`
	name string `uri:"name" binding:"required"`
}

func main() {
	router := gin.Default()
	router.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.Status(404)
			return
		}
		c.JSON(200, gin.H{
			"name": person.name,
			"id":   person.ID,
		})
	})
	router.Run()
}
