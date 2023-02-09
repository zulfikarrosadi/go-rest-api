package note

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type noteHandler struct {
	*service
}

const (
	contentType = "Content-Type"
	appJson     = "application/json"
)

func NewHandler(s *service) *noteHandler {
	return &noteHandler{
		service: s,
	}
}

func (h *noteHandler) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("noteId"))
	response := h.service.Get(r.Context(), id)
	w.Header().Add(contentType, appJson)
	json.NewEncoder(w).Encode(&response)
}

func (h *noteHandler) Query(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response := h.service.Query(r.Context())
	w.Header().Add(contentType, appJson)
	json.NewEncoder(w).Encode(&response)
}

func (h *noteHandler) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	note := &NoteCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(note)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic("data is not valid")
	}

	response := h.service.Create(r.Context(), note)
	w.Header().Add(contentType, appJson)
	json.NewEncoder(w).Encode(&response)
}

func (h *noteHandler) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	note := &NoteUpdateRequest{}
	err := json.NewDecoder(r.Body).Decode(note)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic("data is not valid")
	}

	response := h.service.Update(r.Context(), note)
	w.Header().Add(contentType, appJson)
	json.NewEncoder(w).Encode(response)
}

func (h *noteHandler) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	noteId, err := strconv.Atoi(p.ByName("noteId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic("check your request url")
	}

	response := h.service.Delete(r.Context(), noteId)
	w.Header().Add(contentType, appJson)
	json.NewEncoder(w).Encode(response)
}
