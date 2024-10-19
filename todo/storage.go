package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// InitTodosFile creates the JSON which stores the todos if it does not exist
// and initialize it with an empty array.
func InitTodosFile() error {
	fileInfo, err := os.Stat("todos.json")

	if fileInfo != nil {
		return nil
	}

	// We simply ignore the ErrNotExist and handle other errors
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("error checking todos.json: %w", err)
	}

	f, err := os.Create("todos.json")
	if err != nil {
		return fmt.Errorf("failed to create todos.json: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString("[]\n"); err != nil {
		return fmt.Errorf("failed to write to todos.json: %w", err)
	}

	return nil
}

func ReadTodosFile() ([]Todo, error) {
	b, err := os.ReadFile("todos.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read todos.json: %w", err)
	}

	var todos []Todo

	if err := json.Unmarshal(b, &todos); err != nil {
		return nil, fmt.Errorf("failed to unmarshal todos: %w", err)
	}

	return todos, nil
}

// Save a slice of todo in the file.
// Overwrites all the todos already present in the file.
func SaveTodos(todos []Todo) error {
	todosB, err := json.MarshalIndent(todos, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal todos: %w", err)
	}

	if err := os.WriteFile("todos.json", todosB, 0666); err != nil {
		return fmt.Errorf("failed to write todos to file: %w", err)
	}

	return nil
}
