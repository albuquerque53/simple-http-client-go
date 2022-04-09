package model

import (
	"log"
	"net/http"
)

func CreateRequest(method, uri string) *http.Request {
	request, err := http.NewRequest(method, uri, nil)

	if err != nil {
		log.Fatal(err)
	}

	return request
}
