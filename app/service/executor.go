package service

import (
	"io/ioutil"
	"log"
	"net/http"
)

func ExecuteRequest(simpleRequest *http.Request) string {
	client := &http.Client{}

	response, err := client.Do(simpleRequest)

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
