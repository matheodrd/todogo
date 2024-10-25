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

func (t *Todo) SetTitle(newTitle string) {
	t.Title = newTitle
}

func (t *Todo) SetDescription(newDescription string) {
	t.Description = newDescription
}

func (t Todo) String() string {
	return fmt.Sprintf("%s | %s | %s | %s", t.ID, t.Title, t.Description, t.Status.String())
}

type TodoList struct {
	Todos []*Todo
}

func NewTodoList() (*TodoList, error) {
	todos, err := ReadTodosFile()
	if err != nil {
		return nil, fmt.Errorf("error while reading todos from file: %w", err)
	}

	// must convert todos read from file to a slice of pointers
	todoPtrs := make([]*Todo, len(todos))
	for i := range todos {
		todoPtrs[i] = &todos[i]
	}

	return &TodoList{
		Todos: todoPtrs,
	}, nil
}

func (tl *TodoList) Display() {
	for _, todo := range tl.Todos {
		fmt.Print(todo)
	}
}

func (tl *TodoList) AddTodo(t Todo) {
	tl.Todos = append(tl.Todos, &t)
}

func (tl *TodoList) FindTodo(id uuid.UUID) (*Todo, error) {
	for _, todo := range tl.Todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return nil, fmt.Errorf("cannot find todo with id %s", id)
}

func (tl *TodoList) RemoveTodo(id uuid.UUID) (string, error) {
	for idx, todo := range tl.Todos {
		if todo.ID == id {
			tl.Todos = slices.Delete(tl.Todos, idx, idx+1)
			return todo.Title, nil
		}
	}
	return "", fmt.Errorf("unable to remove todo: cannot find todo with id %s", id)
}

func (tl *TodoList) UpdateTodoStatus(id uuid.UUID, newStatus status) error {
	todo, err := tl.FindTodo(id)
	if err != nil {
		return fmt.Errorf("unable to update todo status: %w", err)
	}
	todo.SetStatus(newStatus)
	return nil
}

func (tl *TodoList) UpdateTodoTitle(id uuid.UUID, newTitle string) error {
	todo, err := tl.FindTodo(id)
	if err != nil {
		return fmt.Errorf("unable to update todo title: %w", err)
	}
	todo.SetTitle(newTitle)
	return nil
}

func (tl *TodoList) UpdateTodoDescription(id uuid.UUID, newDescription string) error {
	todo, err := tl.FindTodo(id)
	if err != nil {
		return fmt.Errorf("unable to update todo description: %w", err)
	}
	todo.SetDescription(newDescription)
	return nil
}
