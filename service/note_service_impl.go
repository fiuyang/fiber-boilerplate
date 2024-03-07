package service

import (
	"context"
	"fiber-boilerplate/data/request"
	"fiber-boilerplate/data/response"
	"fiber-boilerplate/helper"
	"fiber-boilerplate/model"
	"fiber-boilerplate/repository"

	"fiber-boilerplate/exception"
	"github.com/go-playground/validator/v10"
)

type NoteServiceImpl struct {
	NoteRepository repository.NoteRepository
	validate       *validator.Validate
}

func NewNoteServiceImpl(noteRepository repository.NoteRepository, validate *validator.Validate) NoteService {
	return &NoteServiceImpl{
		NoteRepository: noteRepository,
		validate:       validate,
	}
}

// Create implements NoteService
func (service *NoteServiceImpl) Create(ctx context.Context, request request.CreateNoteRequest) {
	err := service.validate.Struct(request)
	helper.ErrorPanic(err)

	dataset := model.Note{
		Content: request.Content,
	}

	service.NoteRepository.Insert(ctx, dataset)
}

// Delete implements NoteService
func (service *NoteServiceImpl) Delete(ctx context.Context, noteId int) {
	service.NoteRepository.Delete(ctx, noteId)
}

// FindAll implements NoteService
func (service *NoteServiceImpl) FindAll(ctx context.Context) []response.NoteResponse {
	result := service.NoteRepository.FindAll(ctx)
	var notes []response.NoteResponse

	for _, value := range result {
		note := response.NoteResponse{
			ID:      value.ID,
			Content: value.Content,
		}

		notes = append(notes, note)
	}

	return notes
}

// FindById implements NoteService
func (service *NoteServiceImpl) FindById(ctx context.Context, noteId int) response.NoteResponse {
	data, err := service.NoteRepository.FindById(ctx, noteId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	response := response.NoteResponse{
		ID:      data.ID,
		Content: data.Content,
	}

	return response
}
