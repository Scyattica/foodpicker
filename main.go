package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	client, ctx, cancel, err := connect("mongodb://192.168.1.193:27017")
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)

	// Ping mongoDB with Ping method
	ping(client, ctx)
}
