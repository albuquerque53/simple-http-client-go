package main

import (
	"fmt"
	"os"
	"simplehttpclient/app/model"
	"simplehttpclient/app/service"
)

func main() {
	httpMethod := os.Args[1]
	target := os.Args[2]

	request := model.CreateRequest(httpMethod, target)
	response := service.ExecuteRequest(request)

	fmt.Println(response)
}
