package handler

import (
	"github.com/gorilla/mux"
)

func (h *handler) MapRoutes() *mux.Router {
	r := mux.NewRouter()

	p := r.PathPrefix("/public").Subrouter()
	p.HandleFunc("/create_thread", h.CreateThread)
	p.HandleFunc("/get_thread", h.GetThread)
	p.HandleFunc("/create_vote", h.CreateVote)
	p.HandleFunc("/use_vote", h.UseVote)
	p.HandleFunc("/end_thread", h.EndThread)

	a := r.PathPrefix("/anon").Subrouter()
	a.HandleFunc("/create_thread", h.CreateAnonThread)
	a.HandleFunc("/get_thread", h.GetAnonThread)
	a.HandleFunc("/create_vote", h.CreateVote)
	a.HandleFunc("/use_vote", h.UseAnonVote)
	a.HandleFunc("/end_thread", h.EndAnonThread)

	return r
}