package main

import (
	"fmt"

	"github.com/Rivalz-ai/proto/client/go/price"
)

func main() {
	fmt.Println("hello proto")
	err := price.NewPriceServiceClient("localhost:50051")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("price service client created")

}
