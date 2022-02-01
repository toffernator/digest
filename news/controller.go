package news

type NewsController struct {
	readers map[string]Reader
}

func NewController(readers map[string]Reader) *NewsController {
	return &NewsController{readers: readers}
}

func (c *NewsController) AllArticles() []Article {
	articles := make([]Article, 0)
	for _, r := range c.readers {
		articles = append(articles, r.Read()...)
	}
	return articles
}

func (c *NewsController) ArticlesFrom(src string) []Article {
	if r, found := c.readers[src]; found {
		return r.Read()
	}

	return make([]Article, 0)
}
