package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/ashwins93/wutplans"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type CreateTodoParams struct {
	ID          string
	Task        string
	CompletedAt sql.NullInt64 `db:"completed_at"`
	CreatedAt   int64         `db:"created_at"`
}

type TodoRow struct {
	ID          string
	Task        string
	CompletedAt sql.NullInt64 `db:"completed_at"`
	CreatedAt   int64         `db:"created_at"`
}

const (
	CREATE_TODO_QUERY = `INSERT INTO TODOS VALUES(:id, :task, :completed_at, :created_at)`
	FIND_TODO_BY_ID   = `SELECT ID, TASK, COMPLETED_AT, CREATED_AT FROM todos`
)

func (db *DB) CreateTodo(ctx context.Context, todo *wutplans.Todo) (*wutplans.Todo, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}

	todo.ID = id
	todo.CreatedAt = time.Now()

	createTodoParams := CreateTodoParams{
		ID:          todo.ID,
		Task:        todo.Task,
		CompletedAt: sql.NullInt64{},
		CreatedAt:   todo.CreatedAt.Unix(),
	}

	_, err = db.db.NamedExecContext(ctx, CREATE_TODO_QUERY, createTodoParams)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (db *DB) FindAllTodos(ctx context.Context) ([]*wutplans.Todo, error) {
	results := make([]TodoRow, 0)

	err := db.db.SelectContext(ctx, &results, FIND_TODO_BY_ID)
	if err != nil {
		return nil, err
	}

	todos := make([]*wutplans.Todo, len(results))
	for i, todoRow := range results {
		todo := new(wutplans.Todo)
		todo.ID = todoRow.ID
		todo.Task = todoRow.Task
		todo.CreatedAt = time.Unix(todoRow.CreatedAt, 0)

		if todoRow.CompletedAt.Valid {
			completedAt := time.Unix(todoRow.CompletedAt.Int64, 0)
			todo.CompletedAt = &completedAt
		}

		todos[i] = todo
	}

	return todos, nil
}
