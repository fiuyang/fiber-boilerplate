package controller

import (
	"context"
	"fiber-boilerplate/data/request"
	"fiber-boilerplate/data/response"
	"fiber-boilerplate/helper"
	"fiber-boilerplate/service"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"fiber-boilerplate/utils"
)

type NoteController struct {
	noteService service.NoteService
}

func NewNoteController(service service.NoteService) *NoteController {
	return &NoteController{
		noteService: service,
	}
}

// Note             godoc
// @Summary			Create note
// @Description		Save note data in Db.
// @Param			note body request.CreateNoteRequest true "Create note"
// @Produce			application/json
// @Tags			note
// @Success			200 {object} response.Response{}
// @Router			/notes [post]
func (controller *NoteController) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	request := request.CreateNoteRequest{}
	err := ctx.BodyParser(&request)
	helper.ErrorPanic(err)

	controller.noteService.Create(c, request)

	webResponse := response.Response{
		Code:    fiber.StatusCreated,
		Status:  "Created",
		Message: "Created Successfully",
		Data:    nil,
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

// Note		        godoc
// @Summary			Delete note
// @Param		    noteId path string true "delete note by id"
// @Description		Remove note data by id.
// @Produce			application/json
// @Tags			note
// @Success			200 {object} response.Response{}
// @Router			/notes/{noteId} [delete]
func (controller *NoteController) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	paramId := ctx.Params("noteId")
	id, err := strconv.Atoi(paramId)
	helper.ErrorPanic(err)
	controller.noteService.Delete(c, id)

	webResponse := response.Response{
		Code:    fiber.StatusOK,
		Status:  "OK",
		Message: "Deleted Successfully",
		Data:    nil,
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

// Note 		        godoc
// @Summary				Get Single note by id.
// @Param				noteId path string true "get note by id"
// @Description			Return the note whoes noteId value mathes id.
// @Produce				application/json
// @Tags				note
// @Success				200 {object} response.Response{}
// @Router				/notes/{noteId} [get]
func (controller *NoteController) FindById(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	paramId := ctx.Params("noteId")
	id, err := strconv.Atoi(paramId)
	helper.ErrorPanic(err)

	data := controller.noteService.FindById(c, id)

	webResponse := response.Response{
		Code:    fiber.StatusOK,
		Status:  "OK",
		Data:    data,
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

// Note             godoc
// @Summary			Get All note.
// @Description		Return list of note.
// @Produce		    application/json
// @Tags			note
// @Success         200 {object} response.Response{}
// @Router			/notes [get]
func (controller *NoteController) FindAll(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	data := controller.noteService.FindAll(c)

	webResponse := response.Response{
		Code:    fiber.StatusOK,
		Status:  "OK",
		Data:    data,
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}