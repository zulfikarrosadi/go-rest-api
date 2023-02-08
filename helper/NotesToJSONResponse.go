package helper

import (
	"github.com/zulfikarrosadi/go-rest-pg/model"
	"github.com/zulfikarrosadi/go-rest-pg/model/web"
)

func NoteToJSONResponse(data []model.Note, errors error) *web.NoteResponse {
	if errors != nil {
		return &web.NoteResponse{
			Status: "failed",
			Error:  errors.Error(),
			Data:   nil,
		}
	}
	notes := []web.Note{}
	for _, v := range data {
		notes = append(notes, web.Note(v))
	}
	return &web.NoteResponse{
		Status: "success",
		Error:  "",
		Data:   notes,
	}
}
