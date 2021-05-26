package Interfaces

type IRepository interface {
	GetAll(entities interface{}) error
	GetByID(entity interface{}, id string) error
	Insert(entities interface{}) error
	Update(entities interface{}) error
	UpdateColumn(entity interface{}) error
	Delete(entity interface{}) error

	Commit()
}
