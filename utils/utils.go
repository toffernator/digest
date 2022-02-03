package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Must(err error) {
	if err != nil {
		log.Fatalf("Failed with err: %v", err)
	}
}

func Get(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	return nil
}
