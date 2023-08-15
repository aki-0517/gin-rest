package main

import (
	"gin-rest/handler"
	"gin-rest/infla"
	"gin-rest/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	d, err := infla.NewDB(100)
	if err != nil {
		fmt.Printf("failed to start server. db setup failed, err = %s", err.Error())
		return
	}
	r := setupRouter(d)
	r.Run(":8000")
}

func setupRouter(d *gorm.DB) *gin.Engine {
	r := gin.Default()

	repository := infla.NewPlayer(d)
	usecase := usecase.NewPlayer(repository)
	handler := handler.NewPlayer(usecase)

	todo := r.Group("/todo")
	{
		todo.POST("", handler.Create)
		todo.GET("", handler.FindAll)
		todo.GET("/:id", handler.Find)
	}
	return r
}