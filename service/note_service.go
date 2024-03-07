package service

import (
	"context"
	"fiber-boilerplate/data/request"
	"fiber-boilerplate/data/response"
)

type NoteService interface {
	Create(ctx context.Context, request request.CreateNoteRequest)
	Delete(ctx context.Context, noteId int)
	FindById(ctx context.Context, noteId int) response.NoteResponse
	FindAll(ctx context.Context) []response.NoteResponse
}
