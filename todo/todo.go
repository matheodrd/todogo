package todo

import "fmt"

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
