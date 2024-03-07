package repository

import (
	"context"
	"errors"
	"fiber-boilerplate/helper"
	"fiber-boilerplate/model"

	"gorm.io/gorm"
)

type NoteRepositoryImpl struct {
	Db *gorm.DB
}

func NewNoteRepositoryImpl(db *gorm.DB) NoteRepository {
	return &NoteRepositoryImpl{Db: db}
}

// Delete implements NoteRepository
func (repo *NoteRepositoryImpl) Delete(ctx context.Context, noteId int) error {
	var note model.Note

	result := repo.Db.WithContext(ctx).Where("id = ?", noteId).Delete(&note)
	if result.Error != nil {
		helper.ErrorPanic(result.Error)
	}

	if result.RowsAffected == 0 {
		return errors.New("pjp is not found")
	}

	return nil

}

// FindAll implements NoteRepository
func (repo *NoteRepositoryImpl) FindAll(ctx context.Context) []model.Note {
	var note []model.Note

	result := repo.Db.WithContext(ctx).Find(&note)
	helper.ErrorPanic(result.Error)
	return note
}

// FindById implements NoteRepository
func (repo *NoteRepositoryImpl) FindById(ctx context.Context, noteId int) (model.Note, error) {
	var note model.Note

	result := repo.Db.WithContext(ctx).First(&note, noteId)
	if result.Error != nil {
		return note, result.Error
	}

	return note, nil
}

// Insert implements NoteRepository
func (repo *NoteRepositoryImpl) Insert(ctx context.Context, note model.Note) {
	result := repo.Db.WithContext(ctx).Create(&note)
	helper.ErrorPanic(result.Error)
}

// Update implements NoteRepository
func (*NoteRepositoryImpl) Update(ctx context.Context, note model.Note) {
	panic("unimplemented")
}
