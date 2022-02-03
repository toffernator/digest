package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/toffernator/digest/news"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	r := gin.Default()

	newsGroup := r.Group("/news")
	{
		newsGroup.GET("", news.AllArticlesHandler)
		newsGroup.GET("/:src", news.ArticlesFromHandler)
	}

	r.Run(":8080")
}
