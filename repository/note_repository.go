package repository

import (
	"context"
	"fiber-boilerplate/model"
)

type NoteRepository interface {
	Insert(ctx context.Context, note model.Note)
	Update(ctx context.Context, note model.Note)
	Delete(ctx context.Context, noteId int) error
	FindById(ctx context.Context, noteId int) (model.Note, error)
	FindAll(ctx context.Context) []model.Note
}
