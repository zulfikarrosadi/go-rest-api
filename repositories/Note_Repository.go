package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zulfikarrosadi/go-rest-pg/helper"
	"github.com/zulfikarrosadi/go-rest-pg/model"
	"github.com/zulfikarrosadi/go-rest-pg/model/web"
)

type NoteRepository struct {
	DB *pgxpool.Pool
}

func NewNoteRepository(db *pgxpool.Pool) *NoteRepository {
	return &NoteRepository{
		DB: db,
	}
}

func (np *NoteRepository) FindAll(ctx context.Context) []model.Note {
	r, err := np.DB.Query(ctx, "SELECT id, title, description FROM notes")
	helper.PanicIfErr(err)
	defer r.Close()

	notes := []model.Note{}
	for r.Next() {
		note := model.Note{}
		err := r.Scan(&note.Id, &note.Title, &note.Description)
		helper.PanicIfErr(err)
		notes = append(notes, note)
	}

	return notes
}

func (np *NoteRepository) Create(ctx context.Context, note web.NoteCreateRequest) (int, error) {
	ct, err := np.DB.Exec(ctx, "INSERT INTO notes (title, description) VALUES ($1,$2)", note.Title, note.Description)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("cannot create note")
	}

	rowsAffected := ct.RowsAffected()
	return int(rowsAffected), nil
}

func (np *NoteRepository) FindByID(ctx context.Context, noteId int) model.Note {
	noteFromDB := np.DB.QueryRow(ctx, "SELECT id, title, description FROM notes WHERE id = $1", noteId)
	note := &model.Note{}
	err := noteFromDB.Scan(&note.Id, &note.Title, &note.Description)
	if err != nil {
		fmt.Println(err)
		return *note
	}
	return *note
}

func (np *NoteRepository) Update(ctx context.Context, note web.NoteCreateRequest, id int) error {
	_, err := np.DB.Exec(ctx, "UPDATE notes SET title = $1, description = $2 WHERE id = $3", note.Title, note.Description, id)
	if err != nil {
		fmt.Println(err)
		return errors.New("cannot update the note")
	}

	return nil
}

func (np *NoteRepository) DeleteByID(ctx context.Context, id int) error {
	_, err := np.DB.Exec(ctx, "DELETE FROM notes WHERE id = $1", id)
	if err != nil {
		fmt.Println(err)
		return errors.New("cannot delete the note")
	}
	return nil
}
