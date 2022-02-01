package news

import (
	"reflect"
	"testing"
	"time"
)

func TestAllArticlesReturnsTwoArticles(t *testing.T) {
	readers := map[string]Reader{
		"stub1": NewStubReader("stub1"),
		"stub2": NewStubReader("stub2"),
	}
	controller := NewController(readers)

	actual := controller.AllArticles()

	expectedOutput := []Article{
		{
			Title:   "Fake Article",
			Summary: "A very awesome article, totally exists. You wish you could read it",
			Author:  "Perry the Platypus",
			Date:    time.Date(2042, time.January, 1, 0, 0, 0, 0, time.UTC),
			Url:     "www.the-best-news-site.org",
		},
		{
			Title:   "Fake Article",
			Summary: "A very awesome article, totally exists. You wish you could read it",
			Author:  "Perry the Platypus",
			Date:    time.Date(2042, time.January, 1, 0, 0, 0, 0, time.UTC),
			Url:     "www.the-best-news-site.org",
		},
	}

	if !reflect.DeepEqual(actual, expectedOutput) {
		t.Errorf("Expected %v but got %v", expectedOutput, actual)
	}
}

func TestArticlesFromGivenStub1ReturnsOneArticle(t *testing.T) {
	readers := map[string]Reader{
		"stub1": NewStubReader("stub1"),
		"stub2": NewStubReader("stub2"),
	}
	controller := NewController(readers)

	actual := controller.ArticlesFrom("stub1")

	expectedOutput := []Article{
		{
			Title:   "Fake Article",
			Summary: "A very awesome article, totally exists. You wish you could read it",
			Author:  "Perry the Platypus",
			Date:    time.Date(2042, time.January, 1, 0, 0, 0, 0, time.UTC),
			Url:     "www.the-best-news-site.org",
		},
	}

	if !reflect.DeepEqual(actual, expectedOutput) {
		t.Errorf("Expected %v but got %v", expectedOutput, actual)
	}
}
