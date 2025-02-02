package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"personal_security/internal/usecase"
)

type Server struct {
	host   string
	port   uint16
	router *gin.Engine
}

func (s *Server) ServeHTTP(_ http.ResponseWriter, _ *http.Request) {

}

type UseCases struct {
	User     *usecase.User
	Event    *usecase.Event
	Contact  *usecase.Contact
	Reminder *usecase.Reminder
}

func NewServer(useCases UseCases, options ...func(*Server)) *Server {
	r := gin.Default()

	setupRouter(r, useCases)

	s := &Server{router: r, host: "0.0.0.0", port: 8080}
	for _, o := range options {
		o(s)
	}

	return s
}

func (s *Server) Run(_ context.Context) error {
	return s.router.Run(fmt.Sprintf("%s:%d", s.host, s.port))
}
