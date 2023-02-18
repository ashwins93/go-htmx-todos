package http

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/ashwins93/wutplans"
	"github.com/labstack/echo/v4"
)

func (s *Server) handleIndex(ctx echo.Context) error {
	todos, err := s.todosService.FindAllTodos(ctx.Request().Context())
	if err != nil {
		log.Printf("Cannot fetch todos %s\n", err.Error())
		return err
	}
	return ctx.Render(http.StatusOK, "index.html", map[string]any{
		"title": "Todos App",
		"todos": todos,
	})
}

func (s *Server) handleAddTodo(ctx echo.Context) error {
	todo := new(wutplans.Todo)

	if err := ctx.Bind(todo); err != nil {
		return err
	}

	_, err := s.todosService.CreateTodo(ctx.Request().Context(), todo)
	if err != nil {
		return err
	}
	todos, err := s.todosService.FindAllTodos(ctx.Request().Context())
	if err != nil {
		return err
	}

	return ctx.Render(http.StatusOK, "todos.html", map[string][]*wutplans.Todo{
		"todos": todos,
	})
}

func (s *Server) handleToggleCompleteTodo(ctx echo.Context) error {
	id := ctx.Param("id")

	if id == "" {
		return errors.New("id not specified")
	}

	currentTime := time.Now()
	_, err := s.todosService.UpdateTodo(ctx.Request().Context(), id, wutplans.Todo{
		CompletedAt: &currentTime,
	})
	if err != nil {
		return err
	}

	todos, err := s.todosService.FindAllTodos(ctx.Request().Context())
	if err != nil {
		return err
	}

	return ctx.Render(http.StatusOK, "todos.html", echo.Map{
		"todos": todos,
	})
}
