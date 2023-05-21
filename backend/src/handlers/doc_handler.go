package handlers

import (
	"net/http"
)

type ApiDocHandler struct {
	fileServer http.Handler
}

func NewApiDocHandler() *ApiDocHandler {
	filesDir := http.Dir("./../spec")
	filesServer := http.FileServer(filesDir)
	return &ApiDocHandler{fileServer: filesServer}
}

func (sh *ApiDocHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sh.fileServer.ServeHTTP(w, r)
}