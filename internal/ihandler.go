package internal

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler interface {
	CreateThread(w http.ResponseWriter, r *http.Request)
	GetThread(w http.ResponseWriter, r *http.Request)
	CreateVote(w http.ResponseWriter, r *http.Request)
	UseVote(w http.ResponseWriter, r *http.Request)
	EndThread(w http.ResponseWriter, r *http.Request)

	CreateAnonThread(w http.ResponseWriter, r *http.Request)
	GetAnonThread(w http.ResponseWriter, r *http.Request)
	UseAnonVote(w http.ResponseWriter, r *http.Request)
	EndAnonThread(w http.ResponseWriter, r *http.Request)

	MapRoutes() *mux.Router
}
