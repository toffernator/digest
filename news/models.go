package news

import "time"

type Article struct {
	Title string     `json:"title"`
	Tags  []string   `json:"tags"`
	Src   string     `json:"src'`
	Date  *time.Time `json:"date"`
	Url   string     `json:"url"`
}
