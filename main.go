package main

import (
	"log"

	"github.com/toffernator/digest/news"
)

func main() {
	readers := map[string]news.Reader{
		news.DR: news.NewDrReader(),
	}

	newsController := news.NewController(readers)

	for _, a := range newsController.AllArticles() {
		log.Println(a)
	}
}
