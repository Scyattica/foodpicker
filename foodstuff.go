package main

import (
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type Foods struct {
	name string `bson:name`
}

func allfoods(w http.ResponseWriter, r *http.Request) {
	client, ctx, cancel, err := connect(getMongoConnectionString())
	if err != nil {
		fmt.Fprintf(w, "MongoDB connection error!")
	}

	collection := client.Database("foodnstuff").Collection("restaurants")
	cursor, colerr := collection.Find(ctx, bson.D{})
	if colerr != nil {
		fmt.Fprintf(w, "Issue trying to find data in the restaurants table!")
	}
	defer close(client, ctx, cancel)
	var restaurants []bson.M
	if err = cursor.All(ctx, &restaurants); err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, restaurants)
}
