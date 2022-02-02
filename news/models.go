package news

import "time"

type Article struct {
	Title   string     `json:"title"`
	Summary string     `json:"summary"`
	Author  string     `json:"author"`
	Date    *time.Time `json:"date"`
	Url     string     `json:"url"`
}
