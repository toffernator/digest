package news

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/toffernator/digest/utils"
)

type Reader func() []Article

func DrReader() []Article {
	resp, err := http.Get(DR_URL)
	utils.Must(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	utils.Must(err)

	var drArticles DrResponse
	err = json.Unmarshal(body, &drArticles)
	utils.Must(err)

	articles := make([]Article, 0)
	for _, a := range drArticles.Data.FrontPage.TopStories {
		article := Article{
			Title:   a.Title,
			Summary: a.Article.Summary,
			Author:  a.Article.Contributions[0].Agent.Name,
			Date:    a.Article.StartDate,
			Url:     a.URL,
		}

		articles = append(articles, article)
	}

	return articles
}
