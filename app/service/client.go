package service

import (
	"io/ioutil"
	"log"
	"net/http"
)

func MakeRequest(method, uri string) string {
	client := &http.Client{}

	request, err := http.NewRequest(method, uri, nil)

	if err != nil {
		log.Fatal("Error")
	}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
