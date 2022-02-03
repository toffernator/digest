package news

import (
	"github.com/mmcdole/gofeed"
	"github.com/toffernator/digest/news/guardian"
)

type Reader func() []Article

func DrReader() []Article {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.dr.dk/nyheder/service/feeds/senestenyt")

	articles := make([]Article, len(feed.Items))
	for i, a := range feed.Items {
		article := Article{
			Title:   a.Title,
			Summary: a.Description,
			Author:  "",
			Date:    a.PublishedParsed,
			Url:     a.Link,
		}

		articles[i] = article
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
			Title:   a.WebTitle,
			Summary: "",
			Author:  "",
			Date:    &a.WebPublicationDate,
			Url:     a.WebURL,
		}
	}

	return articles
}
