package dr

import "github.com/mmcdole/gofeed"

const URL = ""

func Get() (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://www.dr.dk/nyheder/service/feeds/senestenyt")
	if err != nil {
		return nil, err
	}

	return feed, nil
}
