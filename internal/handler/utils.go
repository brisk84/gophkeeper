package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type NilType interface{}

func parseRequest[T any](r *http.Request) (T, error) {
	var data T
	err := json.NewDecoder(r.Body).Decode(&data)
	return data, err
}

// func sendResponse[T any](w http.ResponseWriter, data T, err error) {
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	if err := json.NewEncoder(w).Encode(data); err != nil {
// 		panic(err)
// 	}
// }

func sendResponse(w http.ResponseWriter, data any, statusCode int) {
	w.WriteHeader(statusCode)
	if statusCode == http.StatusNoContent {
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println(err)
	}
}
