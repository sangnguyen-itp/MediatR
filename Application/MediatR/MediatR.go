package MediatR

type IMediatR interface {
	Send(request IRequest) IResponse
}

type IRequest interface {

}

type IResponse interface {

}

type IRequestHandler interface {
	Handle(request IRequest)
}
