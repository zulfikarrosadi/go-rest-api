package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/zulfikarrosadi/go-rest-pg/helper"
	"github.com/zulfikarrosadi/go-rest-pg/model"
	"github.com/zulfikarrosadi/go-rest-pg/model/web"
	"github.com/zulfikarrosadi/go-rest-pg/repositories"
)

type NoteController struct {
	NoteRepository *repositories.NoteRepository
}

func NewNoteController(repository *repositories.NoteRepository) *NoteController {
	return &NoteController{
		NoteRepository: repository,
	}
}

func (nc *NoteController) GetNotes(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	notes := nc.NoteRepository.FindAll(r.Context())
	w.Header().Add("Content-Type", "application/json")
	response := helper.NoteToJSONResponse(notes, nil)

	if len(notes) < 1 {
		fmt.Println("masuk ke sini", len(notes))
		w.WriteHeader(http.StatusNotFound)
		response = helper.NoteToJSONResponse(notes, errors.New("notes not found"))
	}

	e := json.NewEncoder(w)
	err := e.Encode(response)
	helper.PanicIfErr(err)
}

func (nc *NoteController) GetNoteByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	noteIDInParam := p.ByName("noteID")
	noteID, _ := strconv.Atoi(noteIDInParam)
	note := nc.NoteRepository.FindByID(r.Context(), noteID)

	noteResponse := []model.Note{}
	noteResponse = append(noteResponse, note)
	response := helper.NoteToJSONResponse(noteResponse, nil)

	emptyNote := model.Note{}
	if note == emptyNote {
		w.WriteHeader(http.StatusNotFound)
		response = helper.NoteToJSONResponse(noteResponse, errors.New("note not found"))
	}
	e := json.NewEncoder(w)
	err := e.Encode(response)
	helper.PanicIfErr(err)
}

func (nc *NoteController) CreateNote(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var noteReq web.NoteCreateRequest
	var response web.NoteCreateResponse
	err := json.NewDecoder(r.Body).Decode(&noteReq)
	helper.PanicIfErr(err)

	_, err = nc.NoteRepository.Create(r.Context(), noteReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = web.NoteCreateResponse{Status: "failed", Error: err}
	}
	w.WriteHeader(http.StatusCreated)
	response = web.NoteCreateResponse{Status: "success", Error: nil}

	e := json.NewEncoder(w)
	err = e.Encode(response)
	helper.PanicIfErr(err)
}

func (nc *NoteController) UpdateNote(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var noteReq web.NoteCreateRequest
	var response web.NoteCreateResponse
	err := json.NewDecoder(r.Body).Decode(&noteReq)
	helper.PanicIfErr(err)

	noteIDInParam := p.ByName("noteID")
	noteID, _ := strconv.Atoi(noteIDInParam)

	err = nc.NoteRepository.Update(r.Context(), noteReq, noteID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = web.NoteCreateResponse{Status: "failed", Error: err}
	}

	response = web.NoteCreateResponse{Status: "success", Error: nil}
	e := json.NewEncoder(w)
	err = e.Encode(response)
	helper.PanicIfErr(err)
}

func (nc *NoteController) DeleteNoteByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var response web.NoteCreateResponse

	noteIDInParam := p.ByName("noteID")
	noteID, _ := strconv.Atoi(noteIDInParam)

	err := nc.NoteRepository.DeleteByID(r.Context(), noteID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = web.NoteCreateResponse{Status: "failed", Error: err}
		e := json.NewEncoder(w)
		err = e.Encode(response)
		helper.PanicIfErr(err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (nc *NoteController) PanicHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("panicking")
}

func (nc *NoteController) AuthHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("authenticated"))
}
