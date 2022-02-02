package news

import (
	"github.com/mmcdole/gofeed"
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
