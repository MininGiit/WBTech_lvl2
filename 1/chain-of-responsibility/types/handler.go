package types

import "fmt"

type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

func (h *BaseHandler) Handle(request string) {
	if h.next != nil {
		h.next.Handle(request)
	}
}

type ConcreteHandlerA struct {
	BaseHandler
}

func (h *ConcreteHandlerA) Handle(request string) {
	if request == "A" {
		fmt.Println("ConcreteHandlerA handled the request")
	} else {
		fmt.Println("ConcreteHandlerA passed the request")
		h.BaseHandler.Handle(request)
	}
}

type ConcreteHandlerB struct {
	BaseHandler
}

func (h *ConcreteHandlerB) Handle(request string) {
	if request == "B" {
		fmt.Println("ConcreteHandlerB handled the request")
	} else {
		fmt.Println("ConcreteHandlerB passed the request")
		h.BaseHandler.Handle(request)
	}
}

