package main

import (
	"LLd-Test/contracts"
	"LLd-Test/services"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	requestHandler := func(w http.ResponseWriter, req *http.Request) {
		parsedBody := contracts.InputContract{}
		err := json.NewDecoder(req.Body).Decode(&parsedBody)
		if err != nil {
			fmt.Println("Invalid req body")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Invalid req body"))
			return
		}
		err = parsedBody.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			byteErr, _ := json.Marshal(err.Error())
			_, _ = w.Write(byteErr)
			return
		}

		ps := services.NewPathService()
		response, err := ps.GetAllPathsFromGraph(parsedBody)
		if err != nil {
			fmt.Println("Runtime error")
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Encountered runtime error"))
		}

		resBytes, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(resBytes)
		return
	}

	r.HandleFunc("/find-paths", requestHandler).Methods(http.MethodPost)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Encountered err in initiating server: ", err.Error())
	}
}
