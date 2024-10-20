package todo

import (
	"fmt"
	"slices"

	"github.com/google/uuid"
)

type status int

const (
	todo status = iota
	doing
	done
)

func (s status) String() string {
	return [...]string{"to do", "doing", "done"}[s]
}

type Todo struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      status    `json:"status"`
	// CreatedAt   time.Time `json:"created_at"`
}

func NewTodo(title, description string) Todo {
	return Todo{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Status:      todo,
	}
}

func NewTodoWithID(id uuid.UUID, title, description string) Todo {
	return Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      todo,
	}
}

func (t *Todo) SetStatus(newStatus status) {
	t.Status = newStatus
}

type TodoList struct {
	Todos []Todo
}

func NewTodoList() (*TodoList, error) {
	todos, err := ReadTodosFile()
	if err != nil {
		return nil, fmt.Errorf("error while reading todos from file: %w", err)
	}

	return &TodoList{
		Todos: todos,
	}, nil
}

func (tl *TodoList) Display() {
	for _, todo := range tl.Todos {
		fmt.Printf(
			"Tâche n°%d: %s | %s | Statut: %s\n",
			todo.ID, todo.Title, todo.Description, todo.Status.String(),
		)
	}
}

func (tl *TodoList) AddTodo(t Todo) {
	tl.Todos = append(tl.Todos, t)
}

func (tl *TodoList) RemoveTodo(id uuid.UUID) error {
	for idx, todo := range tl.Todos {
		if todo.ID == id {
			tl.Todos = slices.Delete(tl.Todos, idx, idx+1)
			return nil
		}
	}
	return fmt.Errorf("unable to remove todo: cannot find todo with id %d", id)
}

func (tl *TodoList) UpdateTodoStatus(id uuid.UUID, newStatus status) error {
	for idx, todo := range tl.Todos {
		if todo.ID == id {
			t := &tl.Todos[idx]
			t.SetStatus(newStatus)
			return nil
		}
	}
	return fmt.Errorf("unable to update todo status: cannot find todo with id %d", id)
}
