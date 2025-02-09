package http

import (
	"fmt"
	"log"
	"net/http"
	"task12/internal/usecase"
)

type Server struct {
	httpServer	*http.Server 
}

func NewServer(ucEvent usecase.IEvent, port int) *Server {
	handler := NewHandler(ucEvent)
	router := handler.InitRouter()
	addr := fmt.Sprintf(":%d", port)
	server :=  &http.Server{
		Handler:	router,
		Addr:		addr,
	}
	return &Server{
		httpServer: server,
	}
}


func (s *Server) StartServer() {
	log.Println("start server on port:", s.httpServer.Addr)
	s.httpServer.ListenAndServe()
}