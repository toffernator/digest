package news

import "time"

type Reader interface {
	Src() string
	Read() []Article
}

// A stub implementation of the Reader interface
type StubReader struct {
	src string
}

func NewStubReader(src string) *StubReader {
	return &StubReader{src: src}
}

func (r *StubReader) Src() string {
	return r.src
}

func (r *StubReader) Read() []Article {
	return []Article{
		{
			Title:   "Fake Article",
			Summary: "A very awesome article, totally exists. You wish you could read it",
			Author:  "Perry the Platypus",
			Date:    time.Date(2042, time.January, 1, 0, 0, 0, 0, time.UTC),
			Url:     "www.the-best-news-site.org",
		},
	}
}
