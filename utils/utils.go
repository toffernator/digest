package utils

import "log"

func Must(err error) {
	if err != nil {
		log.Fatalf("Failed with err: %v", err)
	}
}
