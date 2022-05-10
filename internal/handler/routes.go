package handler

import (
	"github.com/gorilla/mux"
)

func (h *handler) MapRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/create_thread", h.CreateThread)

	return r
}