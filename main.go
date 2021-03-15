package main

import (
	"fmt"
	"github.com/labstack/echo"
	"main.go/config"
	"main.go/pkg/db/controllers"
	"net/http"
	"os"
)

func main() {
	config.LoadEnvironments()
	fmt.Println("Application started successfully. :)")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to blog APP developed with Golang")
	})
	e.POST("/add-user", controllers.AddUser)
	e.GET("/get-user", controllers.GetUser)
	e.GET("/get-all-user", controllers.GetAllUser)
	e.PUT("/update-user", controllers.UpdateUser)
	e.DELETE("/delete-user", controllers.DeleteUser)
	e.POST("/add-post", controllers.AddPost)
	e.GET("/get-all-post", controllers.GetAllPost)
	e.GET("/get-post", controllers.GetPost)
	e.PUT("/update-post", controllers.UpdatePost)
	//e.DELETE("/delete-post", controllers.DeletePost)
	e.POST("/add-comment", controllers.AddComment)
	//e.PUT("/update-comment", controllers.UpdateComment)
	//e.DELETE("/delete-comment", controllers.DeleteComment)
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
