package news

import (
	"fmt"
)

var readers map[string]Reader = map[string]Reader{
	DR:       DrReader,
	GUARDIAN: GuardianReader,
}

func AllArticles() []Article {
	articles := make([]Article, 0)
	for _, r := range readers {
		articles = append(articles, r()...)
	}

	return articles
}

func ArticlesFrom(src string) ([]Article, error) {
	r, found := readers[src]
	if !found {
		return make([]Article, 0), fmt.Errorf("there is no reader for %s", src)
	}

	return r(), nil
}
