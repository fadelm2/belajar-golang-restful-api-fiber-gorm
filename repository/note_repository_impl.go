package repository

import (
	"errors"
	"gorm.io/gorm"
	"restful-api-gorm-fiber/data/request"
	"restful-api-gorm-fiber/helper"
	"restful-api-gorm-fiber/model"
)

type NoteRepositoryImpl struct {
	Db *gorm.DB
}

func NewNoteRepositoryImpl(Db *gorm.DB) NoteRepository {
	return &NoteRepositoryImpl{Db: Db}
}

func (n *NoteRepositoryImpl) FindById(noteId int) (model.Note, error) {
	var note model.Note
	result := n.Db.First(&note, "id = ?", noteId).Error
	if result == nil {
		return note, nil
	} else {
		return note, errors.New("note is not found")
	}
}
func (n *NoteRepositoryImpl) Save(note model.Note) {
	result := n.Db.Create(&note)
	helper.ErrorPanic(result.Error)
}

func (n *NoteRepositoryImpl) Update(note model.Note) {
	var updateNote = request.UpdateNoteRequest{
		Id:      note.Id,
		Content: note.Content,
	}
	result := n.Db.Model(&note).Where("id = ?", note.Id).Updates(updateNote)
	helper.ErrorPanic(result.Error)
}

func (n *NoteRepositoryImpl) Delete(noteId int) {
	var note model.Note
	result := n.Db.Where("id = ?", noteId).Delete(&note)
	helper.ErrorPanic(result.Error)
}

func (n *NoteRepositoryImpl) FindAll() []model.Note {
	var note []model.Note
	result := n.Db.Find(&note)
	helper.ErrorPanic(result.Error)
	return note
}
