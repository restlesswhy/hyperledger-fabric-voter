package handler

import (
	"encoding/json"
	"fabric-voter/internal"
	"fabric-voter/internal/models"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type ThreadResponse struct {
	ThreadID string `json:"thread_id"`
}

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

	id, err := h.service.CreateThread(params)
	if err != nil {
		http.Error(w, "failed create thread", http.StatusBadRequest)
	}

	httpResp(w, fmt.Sprintf("successfuly created with id: %s", id))
}

func (h *handler) GetThread(w http.ResponseWriter, r *http.Request) {
	threadResp := &ThreadResponse{}
	if err := json.NewDecoder(r.Body).Decode(threadResp); err != nil {
		logrus.Errorf("failed to decode filter from request: %s", err)
		http.Error(w, "failed to decode request object", http.StatusBadRequest)
		return
	}

	thread, err := h.service.GetThread(threadResp.ThreadID)
	if err != nil {
		http.Error(w, "failed to get thread", http.StatusBadRequest)
	}

	httpResp(w, thread)
}

func (h *handler) CreateVote(w http.ResponseWriter, r *http.Request) {
	threadResp := &ThreadResponse{}
	if err := json.NewDecoder(r.Body).Decode(threadResp); err != nil {
		logrus.Errorf("failed to decode filter from request: %s", err)
		http.Error(w, "failed to decode request object", http.StatusBadRequest)
		return
	}

	vote, err := h.service.CreateVote(threadResp.ThreadID)
	if err != nil {
		http.Error(w, "failed create vote", http.StatusBadRequest)
	}

	httpResp(w, vote)
}

func (h *handler) UseVote(w http.ResponseWriter, r *http.Request) {
	voteResp := &models.Vote{}
	if err := json.NewDecoder(r.Body).Decode(voteResp); err != nil {
		logrus.Errorf("failed to decode filter from request: %s", err)
		http.Error(w, "failed to decode request object", http.StatusBadRequest)
		return
	}

	err := h.service.UseVote(voteResp)
	if err != nil {
		http.Error(w, "failed to use vote", http.StatusBadRequest)
	}

	httpResp(w, "successfuly created")
}

func (h *handler) EndThread(w http.ResponseWriter, r *http.Request) {
	threadResp := &ThreadResponse{}
	if err := json.NewDecoder(r.Body).Decode(threadResp); err != nil {
		logrus.Errorf("failed to decode filter from request: %s", err)
		http.Error(w, "failed to decode request object", http.StatusBadRequest)
		return
	}

	err := h.service.EndThread(threadResp.ThreadID)
	if err != nil {
		http.Error(w, "failed create vote", http.StatusBadRequest)
	}

	httpResp(w, "successfuly ended thread")
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
