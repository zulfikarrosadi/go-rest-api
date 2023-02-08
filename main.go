package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zulfikarrosadi/go-rest-pg/controllers"
	"github.com/zulfikarrosadi/go-rest-pg/db"
	"github.com/zulfikarrosadi/go-rest-pg/repositories"
)

func main() {
	noteRepository := repositories.NewNoteRepository(db.GetConnection())
	noteController := controllers.NewNoteController(noteRepository)
	r := httprouter.New()
	r.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-Type", "application/json")
		e := json.NewEncoder(w)
		e.Encode(`{
			"panic": ` + i.(string) + `
		}`)
	}
	r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("Content-Type", "application/json")
		e := json.NewEncoder(w)
		type response struct {
			Status string `json:"status"`
			Errors string `json:"errors"`
		}
		responses := &response{
			Status: "fail",
			Errors: "request url not found",
		}
		e.Encode(responses)
	})

	r.GET("/notes", noteController.GetNotes)
	r.POST("/notes", noteController.CreateNote)
	r.GET("/panic", noteController.PanicHandler)
	r.GET("/notes/:noteID", noteController.GetNoteByID)
	r.PUT("/notes/:noteID", noteController.UpdateNote)
	r.DELETE("/notes/:noteID", noteController.DeleteNoteByID)
	r.POST("/guarded/notes", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		Auth(w, r, p, noteController.AuthHandler)
	})

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: &LogMiddleware{Handler: r},
	}
	server.ListenAndServe()
}

type LogMiddleware struct {
	http.Handler
}

func (m *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("log middleware")
	m.Handler.ServeHTTP(w, r)
}

type AuthHandler func(http.ResponseWriter, *http.Request, httprouter.Params)

func Auth(w http.ResponseWriter, r *http.Request, p httprouter.Params, f AuthHandler) {
	authToken := r.Header.Get("X-TOKEN")
	if authToken == "auth-token" {
		f(w, r, p)
	}
}
