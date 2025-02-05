package handlers

import (
	"LLd-Test/contracts"
	"LLd-Test/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func PathFinderHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
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
}
