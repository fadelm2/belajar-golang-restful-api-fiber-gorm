package service

import (
	"github.com/go-playground/validator/v10"
	"log"
	"restful-api-gorm-fiber/data/request"
	"restful-api-gorm-fiber/data/response"
	"restful-api-gorm-fiber/helper"
	"restful-api-gorm-fiber/model"
	"restful-api-gorm-fiber/repository"
)

type NoteServiceImpl struct {
	NoteRepository repository.NoteRepository
	validate       *validator.Validate
}

func NewNoteServiceImpl(noteRepository repository.NoteRepository, validate *validator.Validate) *NoteServiceImpl {
	return &NoteServiceImpl{
		NoteRepository: noteRepository,
		validate:       validate,
	}
}

func (n *NoteServiceImpl) Create(note request.CreateNoteRequest) {
	err := n.validate.Struct(note)
	helper.ErrorPanic(err)
	noteModel := model.Note{
		Content: note.Content,
	}
	n.NoteRepository.Save(noteModel)
}

func (n *NoteServiceImpl) Update(note request.UpdateNoteRequest) {
	noteData, err := n.NoteRepository.FindById(note.Id)
	helper.ErrorPanic(err)
	noteData.Content = note.Content
	n.NoteRepository.Update(noteData)
}

func (n *NoteServiceImpl) Delete(noteId int) {
	n.NoteRepository.Delete(noteId)
}

func (n *NoteServiceImpl) FindById(noteId int) (response.NoteResponse, error) {
	err := n.validate.Var(noteId, "number")
	if err != nil {
		log.Println(err)
	}
	noteData, err := n.NoteRepository.FindById(noteId)
	if err != nil {
		log.Println(err)
	}
	noteResponse := response.NoteResponse{
		Id:      noteData.Id,
		Content: noteData.Content,
	}

	return noteResponse, err
}

func (n *NoteServiceImpl) FindAll() []response.NoteResponse {
	result := n.NoteRepository.FindAll()
	var notes []response.NoteResponse

	for _, value := range result {
		note := response.NoteResponse{
			Id:      value.Id,
			Content: value.Content,
		}
		notes = append(notes, note)
	}
	return notes
}
