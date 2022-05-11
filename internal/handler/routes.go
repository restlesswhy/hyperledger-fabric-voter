package handler

import (
	"github.com/gorilla/mux"
)

func (h *handler) MapRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/create_thread", h.CreateThread)
	r.HandleFunc("/get_thread", h.GetThread)
	r.HandleFunc("/create_vote", h.CreateVote)
	r.HandleFunc("/use_vote", h.UseVote)
	r.HandleFunc("/end_thread", h.EndThread)

	return r
}