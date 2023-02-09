package note

import (
	"context"

	"github.com/zulfikarrosadi/go-rest-pg/internal/entity"
)

type Service interface {
}

type service struct {
	*NoteRepository
}

func NewService(repo *NoteRepository) *service {
	return &service{
		NoteRepository: repo,
	}
}

type Note struct {
	Note []entity.Note `json:"note"`
}

type NoteCreateRequest struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type NoteUpdateRequest struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Response struct {
	Success bool     `json:"success"`
	Data    any      `json:"data"`
	Error   []string `json:"error"`
}

func (s *service) Query(ctx context.Context) *Response {
	notes, err := s.NoteRepository.Query(ctx)
	if err != nil {
		return &Response{
			Success: false,
			Data:    nil,
			Error:   []string{err.Error()},
		}
	}

	return &Response{
		Success: true,
		Data:    &Note{notes},
		Error:   nil,
	}
}

func (s *service) Get(ctx context.Context, noteId int) *Response {
	n, err := s.NoteRepository.Get(ctx, noteId)
	if err != nil {
		return &Response{
			Success: false,
			Data:    nil,
			Error:   []string{err.Error()},
		}
	}

	return &Response{
		Success: true,
		Data:    &Note{[]entity.Note{n}},
		Error:   nil,
	}
}

func (s *service) Create(ctx context.Context, note *NoteCreateRequest) *Response {
	err := s.NoteRepository.Create(ctx, entity.Note(*note))
	if err != nil {
		return &Response{
			Success: false,
			Data:    nil,
			Error:   []string{err.Error()},
		}
	}

	return &Response{
		Success: true,
		Data:    nil,
		Error:   nil,
	}
}

func (s *service) Update(ctx context.Context, note *NoteUpdateRequest) *Response {
	err := s.NoteRepository.Update(ctx, entity.Note(*note))
	if err != nil {
		return &Response{
			Success: false,
			Data:    nil,
			Error:   []string{err.Error()},
		}
	}

	return &Response{
		Success: true,
		Data:    nil,
		Error:   nil,
	}
}

func (s *service) Delete(ctx context.Context, noteId int) *Response {
	err := s.NoteRepository.Delete(ctx, noteId)
	if err != nil {
		return &Response{
			Success: false,
			Data:    nil,
			Error:   []string{err.Error()},
		}
	}

	return &Response{
		Success: true,
		Data:    nil,
		Error:   nil,
	}
}
