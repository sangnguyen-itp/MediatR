package Repositories

import (
	"gorm.io/gorm"
	"mediatR/Application/Interfaces"
)

type Repository struct {
	_context *gorm.DB
}

func NewGenericRepository(context *gorm.DB) Repository {
	return Repository{_context: context}
}

func (r Repository) GetAll(entities interface{}) error {
	return r._context.Find(entities).Error
}

func (r Repository) GetByID(entity interface{}, id string) error {
	return r._context.First(&entity, "id = ?", id).Error
}

func (r Repository) Insert(entities interface{}) error {
	return r._context.Create(&entities).Error
}

func (r Repository) Update(entities interface{}) error {
	return r._context.Create(&entities).Error
}

func (r Repository) UpdateColumn(entity interface{}) error {
	return r._context.UpdateColumns(entity).Error
}

func (r Repository) Delete(entity interface{}) error {
	return r._context.Unscoped().Delete(entity).Error
}

func (r Repository) Commit() {
	r._context.Commit()
}

var _ Interfaces.IRepository = &Repository{}
