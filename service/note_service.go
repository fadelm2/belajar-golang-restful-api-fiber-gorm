package service

import (
	"restful-api-gorm-fiber/data/request"
	"restful-api-gorm-fiber/data/response"
)

type NoteService interface {
	Create(note request.CreateNoteRequest)
	Update(note request.UpdateNoteRequest)
	Delete(noteId int)
	FindById(noteId int) (response.NoteResponse, error)
	FindAll() []response.NoteResponse
}
