package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	//Let's all set some config stuff!

	fmt.Print(getMongoConnectionString())
	//client, ctx, cancel, err := connect(getMongoConnectionString())
	//if err != nil {
	//	panic(err)
	//}

	//defer close(client, ctx, cancel)

	// Ping mongoDB with Ping method
	//ping(client, ctx)

	handleRequests()
}
