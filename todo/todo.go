package todo

import (
	"fmt"
	"math/rand/v2"
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
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      status `json:"status"`
	// CreatedAt   time.Time `json:"created_at"`
}

func NewTodo(title, description string) Todo {
	return Todo{
		ID:          rand.Int(), // TODO: Replace this crap by a true unique ID i.e. UUID
		Title:       title,
		Description: description,
		Status:      todo,
	}
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
