package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/maicodsantos/newProjectGoweb/cmd/server/handler"
	"github.com/maicodsantos/newProjectGoweb/internal/users"
)

func main() {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("error ao carregar o arquivo .env")
	}
	usuario := os.Getenv("MY_USER")
	password := os.Getenv("MY_PASS")

	log.Println("User: ", usuario)
	log.Println("Password", password)

	repo := users.NewRepository()     // Criação da instância Repository
	service := users.NewService(repo) // Criação da instância Service
	u := handler.NewUser(service)     // Criação do Controller

	r := gin.Default()
	pr := r.Group("/users")
	pr.POST("/post", u.Create())
	pr.GET("/get", u.GetAll())
	pr.PUT("/:id", u.Update())
	pr.PATCH("/:id", u.UpdateName())
	pr.DELETE("/:id", u.Delete())
	r.Run()
}
