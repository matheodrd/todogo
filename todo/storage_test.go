package todo

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
)

// Helper function which creates a temporary directory
// , cd into it and returns the path string of the future JSON file.
func setupTestFilePath(t *testing.T) string {
	t.Helper() // makes it a testing helper function
	tempDir := t.TempDir()
	tempFilePath := filepath.Join(tempDir, "todos.json")

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("failed to cd into the test directory: %v", err)
	}

	t.Cleanup(func() {
		os.Remove(tempFilePath)
	})

	return tempFilePath
}

func TestInitTodosFile(t *testing.T) {
	tempFilePath := setupTestFilePath(t)

	if err := InitTodosFile(); err != nil {
		t.Fatalf("Failed to initialize todos file: %v", err)
	}

	if _, err := os.Stat(tempFilePath); os.IsNotExist(err) {
		t.Fatalf("expected `todos.json` to be created, but it does not exist")
	}
}

func TestSaveTodos(t *testing.T) {
	tempFilePath := setupTestFilePath(t)

	todos := []Todo{
		{ID: uuid.New(), Title: "First todo", Description: "First description", Status: 0},
		{ID: uuid.New(), Title: "Second todo", Description: "Second description", Status: 2},
	}

	err := SaveTodos(todos)
	if err != nil {
		t.Fatalf("Failed to save todos: %v", err)
	}

	content, err := os.ReadFile(tempFilePath)
	if err != nil {
		t.Fatalf("Failed to read file after saving todos: %v", err)
	}

	var result []Todo
	if err := json.Unmarshal(content, &result); err != nil {
		t.Fatalf("Failed to unmarshal todos from file: %v", err)
	}

	if len(result) != len(todos) {
		t.Errorf("expected %d todos, got %d", len(todos), len(result))
	}

	for idx, todo := range result {
		if todo.Title != todos[idx].Title ||
			todo.Description != todos[idx].Description ||
			todo.Status != todos[idx].Status {
			t.Errorf("mismatch in todo at index %d: expected %v, got %v", idx, todos[idx], todo)
		}
	}
}

func TestReadTodosFile(t *testing.T) {
	setupTestFilePath(t)

	todos := []Todo{
		{ID: uuid.New(), Title: "First todo", Description: "First description", Status: 0},
		{ID: uuid.New(), Title: "Second todo", Description: "Second description", Status: 2},
	}

	err := SaveTodos(todos)
	if err != nil {
		t.Fatalf("Failed to save todos: %v", err)
	}

	result, err := ReadTodosFile()
	if err != nil {
		t.Fatalf("Failed to read todos file: %v", err)
	}

	if len(result) != len(todos) {
		t.Errorf("expected %d todos, got %d", len(todos), len(result))
	}

	for idx, todo := range result {
		if todo.Title != todos[idx].Title ||
			todo.Description != todos[idx].Description ||
			todo.Status != todos[idx].Status {
			t.Errorf("mismatch in todo at index %d: expected %v, got %v", idx, todos[idx], todo)
		}
	}
}
