package main

import (
	"fmt"
	"simplehttpclient/app/service"
)

func main() {
	response := service.MakeRequest("GET", "https://duck.com")

	fmt.Println(response)
}
