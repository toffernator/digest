package news

import (
	"github.com/toffernator/digest/news/dr"
	"github.com/toffernator/digest/news/guardian"
)

type Reader func() []Article

func DrReader() []Article {
	resp, err := dr.Get()
	if err != nil {
		return make([]Article, 0)
	}

	articles := make([]Article, len(resp.Items))
	for i, a := range resp.Items {
		articles[i] = Article{
			Title: a.Title,
			Tags:  []string{"Latest News"},
			Src:   "DR",
			Date:  a.PublishedParsed,
			Url:   a.Link,
		}
	}

	return articles
}

func GuardianReader() []Article {
	resp, err := guardian.Get()
	if err != nil {
		return make([]Article, 0)
	}

	articles := make([]Article, len(resp.Data.Results))
	for i, a := range resp.Data.Results {
		articles[i] = Article{
			Title: a.WebTitle,
			Tags:  []string{a.SectionName},
			Src:   "The Guardian",
			Date:  &a.WebPublicationDate,
			Url:   a.WebURL,
		}
	}

	return articles
}
