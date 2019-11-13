package handler

import "net/http"

type MethodHandler map[string]http.Handler
func (mh MethodHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	if h, ok := mh[r.Method]; ok {
		h.ServeHTTP(w, r)
		return
	}
	http.Error(w, "method not allowed.", http.StatusMethodNotAllowed)
}
