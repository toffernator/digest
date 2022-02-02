package news_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toffernator/digest/news"
)

func TestAllArticlesReturnsArticles(t *testing.T) {
	actual := news.AllArticles()

	assert.NotEmpty(t, actual, "Should not be empty")
}

func TestArticlesFromStubReturns1Article(t *testing.T) {
	actual, err := news.ArticlesFrom("dr")

	assert.NoError(t, err, "Should not fail")
	assert.NotEmpty(t, actual, "Should not be empty")
}
