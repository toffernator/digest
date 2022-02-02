package main

import (
	"github.com/gin-gonic/gin"
	"github.com/toffernator/digest/news"
)

func main() {
	r := gin.Default()

	newsGroup := r.Group("/news")
	{
		newsGroup.GET("/", news.AllArticlesHandler)
		newsGroup.GET("/:src", news.ArticlesFromHandler)
	}

	r.Run(":5000")
}
