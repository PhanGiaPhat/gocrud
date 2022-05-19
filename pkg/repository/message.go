package repository

import (
	"github.com/PhanGiaPhat/gocrud/pkg/model"
	"gorm.io/gorm"
)

type MessageRepository interface {
	Create(request model.Message) (*model.Message, error)
	GetByID(id uint) (*model.Message, error)
	List(offset int, size int) ([]model.Message, error)
	Update(id uint, request model.Message) (*model.Message, error)
	Delete(id uint) (*model.Message, error)
}

type messageRepo struct {
	db *gorm.DB
}

func (r *messageRepo) Create(w model.Message) (*model.Message, error) {
	if err := r.db.Create(&w).Error; err != nil {
		return nil, err
	}
	return &w, nil
}

func (r *messageRepo) GetByID(id uint) (*model.Message, error) {
	w := model.Message{
		Model: gorm.Model{ID: id},
	}
	if err := r.db.First(&w).Error; err != nil {
		return nil, err
	}
	return &w, nil
}

func (r *messageRepo) Delete(id uint) (*model.Message, error) {
	w := model.Message{
		Model: gorm.Model{ID: id},
	}
	if err := r.db.Delete(&w).Error; err != nil {
		return nil, err
	}
	return &w, nil
}

func (r *messageRepo) List(offset int, size int) ([]model.Message, error) {
	var ws []model.Message
	paginate := func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(size)
	}
	if err := r.db.Scopes(paginate).Find(&ws).Error; err != nil {
		return nil, err
	}
	return ws, nil
}

func (r *messageRepo) Update(id uint, w model.Message) (*model.Message, error) {
	m := model.Message{
		Model: gorm.Model{ID: id},
	}
	if err := r.db.First(&m).Error; err != nil {
		return nil, err
	}
	m.Type = w.Type
	m.Message = w.Message
	if err := r.db.Save(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func NewMessage(db *gorm.DB) MessageRepository {
	return &messageRepo{db: db}
}
