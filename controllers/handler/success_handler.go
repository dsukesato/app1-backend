package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func Success(w http.ResponseWriter, response interface{}) {
	data, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(w, "marshal error", http.StatusUnauthorized)
		return
	}
	w.Write(data)
}
