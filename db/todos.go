package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/ashwins93/wutplans"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

const (
	CREATE_TODO_QUERY = `INSERT INTO TODOS VALUES(:id, :task, :completed_at, :created_at)`
	FIND_ALL_TODOS    = `SELECT id, task, completed_at, created_at FROM todos`
	FIND_TODO_BY_ID   = `SELECT id, task, completed_at, created_at FROM todos WHERE id = ?`
	UPDATE_TODO       = `UPDATE todos SET task = :task, completed_at = :completed_at WHERE ID = :id`
	DELETE_TODO       = `DELETE FROM todos WHERE id = ?`
)

type TodoRow struct {
	ID          string
	Task        string
	CompletedAt sql.NullInt64 `db:"completed_at"`
	CreatedAt   int64         `db:"created_at"`
}

func (todoRow TodoRow) toDomainModel() *wutplans.Todo {
	todo := new(wutplans.Todo)
	todo.ID = todoRow.ID
	todo.Task = todoRow.Task
	todo.CreatedAt = time.Unix(todoRow.CreatedAt, 0)

	if todoRow.CompletedAt.Valid {
		completedAt := time.Unix(todoRow.CompletedAt.Int64, 0)
		todo.CompletedAt = &completedAt
	}

	return todo
}

func fromDomainModel(todo *wutplans.Todo) TodoRow {
	todoRow := TodoRow{
		ID:          todo.ID,
		Task:        todo.Task,
		CompletedAt: sql.NullInt64{},
		CreatedAt:   todo.CreatedAt.Unix(),
	}

	if todo.CompletedAt != nil {
		todoRow.CompletedAt.Valid = true
		todoRow.CompletedAt.Int64 = todo.CompletedAt.Unix()
	}

	return todoRow
}

func (db *DB) CreateTodo(ctx context.Context, todo *wutplans.Todo) (*wutplans.Todo, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}

	todo.ID = id
	todo.CreatedAt = time.Now()

	newTodoRow := fromDomainModel(todo)

	_, err = db.db.NamedExecContext(ctx, CREATE_TODO_QUERY, newTodoRow)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (db *DB) FindAllTodos(ctx context.Context) ([]*wutplans.Todo, error) {
	results := make([]TodoRow, 0)

	err := db.db.SelectContext(ctx, &results, FIND_ALL_TODOS)
	if err != nil {
		return nil, err
	}

	todos := make([]*wutplans.Todo, len(results))
	for i, todoRow := range results {
		todo := todoRow.toDomainModel()
		todos[i] = todo
	}

	return todos, nil
}

func (db *DB) FindTodoById(ctx context.Context, id string) (*wutplans.Todo, error) {
	todoRow := new(TodoRow)
	err := db.db.GetContext(ctx, todoRow, FIND_TODO_BY_ID, id)
	if err != nil {
		return nil, err
	}

	todo := todoRow.toDomainModel()

	return todo, nil
}

func (db *DB) UpdateTodo(ctx context.Context, id string, updatedTodo wutplans.Todo) (*wutplans.Todo, error) {
	updatedTodoRow := new(TodoRow)

	err := db.db.GetContext(ctx, updatedTodoRow, FIND_TODO_BY_ID, id)
	if err != nil {
		return nil, err
	}

	if updatedTodo.Task != "" {
		updatedTodoRow.Task = updatedTodo.Task
	}

	if updatedTodo.CompletedAt != nil {
		updatedTodoRow.CompletedAt = sql.NullInt64{
			Valid: true,
			Int64: updatedTodo.CompletedAt.Unix(),
		}
	}

	_, err = db.db.NamedExecContext(ctx, UPDATE_TODO, updatedTodoRow)
	if err != nil {
		return nil, err
	}

	todo := updatedTodoRow.toDomainModel()

	return todo, nil
}

func (db *DB) DeleteTodo(ctx context.Context, id string) (*wutplans.Todo, error) {
	todoRow := new(TodoRow)

	err := db.db.GetContext(ctx, todoRow, FIND_TODO_BY_ID, id)
	if err != nil {
		return nil, err
	}

	deletedTodo := todoRow.toDomainModel()

	_, err = db.db.ExecContext(ctx, DELETE_TODO, id)
	if err != nil {
		return nil, err
	}

	return deletedTodo, nil
}
