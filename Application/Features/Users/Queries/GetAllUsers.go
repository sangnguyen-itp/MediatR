package Queries

import "mediatR/Application/MediatR"

type GetAllUser struct {}

func (g GetAllUser) Handle(request MediatR.IRequest) {

}

var _ MediatR.IRequestHandler = &GetAllUser{}