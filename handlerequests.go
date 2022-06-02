package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/mdbhc", mongo_healthcheck)
	myRouter.HandleFunc("/allfoods", allfoods)
    // finally, instead of passing in nil, we want
    // to pass in our newly created router as the second
    // argument
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}