package handler

import (
	"net/http"
)

type IndexHandler struct{}

func (ih *IndexHandler) LookinHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/" {
		http.NotFound(w, r)
		return
	}
	Success(w, IndexResponse{Text:"Welcome to Lookin!"})

	//fmt.Fprint(w, "Welcome to Lookin!")
}

type IndexResponse struct {
	Text string `json:"text"`
}
