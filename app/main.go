package main

import (
	"fmt"
	"os"
	"simplehttpclient/app/service"
)

func main() {
	httpMethod := os.Args[1]
	target := os.Args[2]

	response := service.MakeRequest(httpMethod, target)

	fmt.Println(response)
}
