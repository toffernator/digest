package news

import "time"

type Article struct {
	Title   string
	Summary string
	Author  string
	Date    time.Time
	Url     string
}
