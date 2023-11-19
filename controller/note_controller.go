package controller

import (
	"github.com/gofiber/fiber/v2"
	"restful-api-gorm-fiber/data/request"
	"restful-api-gorm-fiber/data/response"
	"restful-api-gorm-fiber/helper"
	"restful-api-gorm-fiber/service"
	"strconv"
)

type NoteController struct {
	noteService service.NoteService
}

func NewNoteController(service service.NoteService) *NoteController {
	return &NoteController{noteService: service}
}

func (controller *NoteController) Create(ctx *fiber.Ctx) error {
	createNoteRequest := request.CreateNoteRequest{}
	err := ctx.BodyParser(&createNoteRequest)
	helper.ErrorPanic(err)

	controller.noteService.Create(createNoteRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully created note data!",
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *NoteController) Update(ctx *fiber.Ctx) error {
	updateNoteRequest := request.UpdateNoteRequest{}
	err := ctx.BodyParser(&updateNoteRequest)
	helper.ErrorPanic(err)

	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	updateNoteRequest.Id = id
	controller.noteService.Update(updateNoteRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully updated notes data!",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *NoteController) Delete(ctx *fiber.Ctx) error {
	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)
	controller.noteService.Delete(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully deleted notes data!",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)

}
func (controller *NoteController) FindById(ctx *fiber.Ctx) error {
	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	noteResponse := controller.noteService.FindById(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully get notes data!",
		Data:    noteResponse,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}
func (controller *NoteController) FindAll(ctx *fiber.Ctx) error {
	noteResponse := controller.noteService.FindAll()
	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully get notes data!",
		Data:    noteResponse,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)

}
