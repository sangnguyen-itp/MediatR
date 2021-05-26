package Repositories

import (
	"mediatR/Application/Interfaces"
	"mediatR/Infrastructures/Database"
)

type UnitOfWork struct {
	Users Interfaces.IUserRepository
}

var _UnitOfWork *UnitOfWork

func NewUnitOfWork() {
	baseRepository := NewGenericRepository(Database.GetDB())
	_UnitOfWork = &UnitOfWork{
		Users: NewUserRepository(&baseRepository),
	}
}
