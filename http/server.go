package http

import (
	"html/template"

	"github.com/ashwins93/wutplans"

	"github.com/labstack/echo/v4"
)

type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

type Server struct {
	todosService wutplans.TodoService
}

func New(todosService wutplans.TodoService) *Server {
	return &Server{todosService}
}

func (s *Server) SetupRoutes(r EchoRouter) {
	r.GET("/", s.handleIndex)
	r.POST("/todos", s.handleAddTodo)
	r.PUT("/todos/:id", s.handleToggleCompleteTodo)
}

func (s *Server) GetTemplateRenderer(templates *template.Template) *Template {
	return &Template{templates}
}
