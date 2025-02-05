package main

import (
	"LLd-Test/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/find-paths", handlers.PathFinderHandler()).Methods(http.MethodPost)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Encountered err in initiating server: ", err.Error())
	}
}
