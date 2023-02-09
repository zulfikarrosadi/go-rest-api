package note

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zulfikarrosadi/go-rest-pg/internal/entity"
	"github.com/zulfikarrosadi/go-rest-pg/pkg/exception"
)

type Repository interface {
	Get(ctx context.Context, noteId int) (entity.Note, error)
	Query(ctx context.Context) ([]entity.Note, error)
	Create(ctx context.Context, note entity.Note) error
	Update(ctx context.Context, note entity.Note) error
	Delete(ctx context.Context, noteId int) error
}

type NoteRepository struct {
	DB *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *NoteRepository {
	return &NoteRepository{
		DB: db,
	}
}

func (np *NoteRepository) Get(ctx context.Context, noteId int) (entity.Note, error) {
	noteFromDB := np.DB.QueryRow(ctx, "SELECT id, title, description FROM notes WHERE id = $1", noteId)
	note := entity.Note{}
	err := noteFromDB.Scan(&note.Id, &note.Title, &note.Description)
	if err != nil {
		fmt.Println(err)
		return note, err
	}
	return note, nil
}

func (np *NoteRepository) Query(ctx context.Context) ([]entity.Note, error) {
	r, err := np.DB.Query(ctx, "SELECT id, title, description FROM notes")
	exception.PanicIfErr(err)
	defer r.Close()

	notes := []entity.Note{}
	for r.Next() {
		note := entity.Note{}
		err := r.Scan(&note.Id, &note.Title, &note.Description)
		if err != nil {
			return nil, errors.New("notes not found")
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (np *NoteRepository) Create(ctx context.Context, note entity.Note) error {
	_, err := np.DB.Exec(ctx, "INSERT INTO notes (title, description) VALUES ($1,$2)", note.Title, note.Description)
	if err != nil {
		fmt.Println(err)
		return errors.New("cannot create note")
	}

	return nil
}

func (np *NoteRepository) Update(ctx context.Context, note entity.Note) error {
	_, err := np.DB.Exec(ctx, "UPDATE notes SET title = $1, description = $2 WHERE id = $3", note.Title, note.Description, note.Id)
	if err != nil {
		fmt.Println(err)
		return errors.New("cannot update the note")
	}

	return nil
}

func (np *NoteRepository) Delete(ctx context.Context, id int) error {
	_, err := np.DB.Exec(ctx, "DELETE FROM notes WHERE id = $1", id)
	if err != nil {
		fmt.Println(err)
		return errors.New("cannot delete the note")
	}
	return nil
}
