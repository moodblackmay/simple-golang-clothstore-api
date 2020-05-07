package main

import (
	"github.com/gin-gonic/gin"
	"simple-golang-clothstore-api/controllers"
	"simple-golang-clothstore-api/models"
)

func main() {
	s := gin.Default()
	db := models.SetupDB()

	//Provide db to controllers
	s.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	s.GET("/clothes", controllers.FindClothes)
	s.GET("/clothes/:id", controllers.FindCloth)
	s.POST("/clothes", controllers.CreateCloth)
	s.PATCH("/clothes/:id", controllers.UpdateCloth)
	s.DELETE("/clothes/:id", controllers.DeleteCloth)

	_ = s.Run(":8080")
}
