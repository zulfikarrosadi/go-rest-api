package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zulfikarrosadi/go-rest-pg/internal/note"
	"github.com/zulfikarrosadi/go-rest-pg/pkg/db"
)

func main() {
	dbConn := db.GetConnection()
	noteHandler := note.NewHandler(note.NewService(&note.NoteRepository{DB: dbConn}))
	r := httprouter.New()

	r.GET("/notes", noteHandler.Query)
	r.GET("/notes/:noteId", noteHandler.Get)
	r.POST("/notes", noteHandler.Create)
	r.PUT("/notes/:noteId", noteHandler.Update)
	r.DELETE("/notes/:noteId", noteHandler.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: r,
	}
	server.ListenAndServe()
}
