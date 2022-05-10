package handler

import (
	"encoding/json"
	"fabric-voter/internal"
	"fabric-voter/internal/models"
	"net/http"

	"github.com/sirupsen/logrus"
)

type handler struct {
	service internal.Service
}

func NewHandler(service internal.Service) internal.Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateThread(w http.ResponseWriter, r *http.Request) {
	params := &models.ThreadParams{}
	if err := json.NewDecoder(r.Body).Decode(params); err != nil {
		logrus.Errorf("failed to decode filter from request: %s", err)
		http.Error(w, "failed to decode request object", http.StatusBadRequest)
		return
	}

	err := h.service.CreateThread(params)
	if err != nil {
		http.Error(w, "failed create thread", http.StatusBadRequest)
	}    
}

func (h *handler) GetThread(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) CreateVote(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) UseVote(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) EndThread(w http.ResponseWriter, r *http.Request) {

}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func httpResp(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	resp := &Response{
		Status: http.StatusOK,
		Data:   data,
	}

	respBytes, _ := json.Marshal(resp)

	w.Write(respBytes)
}
