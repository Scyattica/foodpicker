package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {
	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

// This is a user defined method that returns mongo.Client,
// context.Context, context.CancelFunc and error.
// mongo.Client will be used for further database operation.
// context.Context will be used set deadlines for process.
// context.CancelFunc will be used to cancel context and
// resource associated with it.

func connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

// This is a user defined method that accepts
// mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func ping(client *mongo.Client, ctx context.Context) string {

	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return "connection failed!"
	}
	return "connected successfully"
}

func getMongoConnectionString() string {
	var mongoip, mongoport, mongouser, mongopass string
	if os.Getenv("MONGO_IP") != "" {
		mongoip = os.Getenv("MONGO_IP")
	} else {
		mongoip = "192.168.1.193"
	}

	if os.Getenv("MONGO_PORT") != "" {
		mongoport = os.Getenv("MONGO_PORT")
	} else {
		mongoport = "27017"
	}

	if os.Getenv("MONGO_USER") != "" {
		mongouser = os.Getenv("MONGO_USER")
	} else {
		mongouser = "root"
	}

	if os.Getenv("MONGO_PASS") != "" {
		mongopass = os.Getenv("MONGO_PASS")
	} else {
		panic("no mongo creds specified!")
	}
	return fmt.Sprintf("mongodb://%s:%s@%s:%s", mongouser, mongopass, mongoip, mongoport)
}

func mongo_healthcheck(w http.ResponseWriter, r *http.Request) {
	client, ctx, cancel, err := connect(getMongoConnectionString())
	if err != nil {
		fmt.Fprintf(w, "MongoDB connection error!")
	}

	defer close(client, ctx, cancel)
	fmt.Fprint(w, ping(client, ctx))
}
