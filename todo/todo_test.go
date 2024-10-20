package todo

import (
	"testing"

	"github.com/google/uuid"
)

func TestAddTodoWithUUID(t *testing.T) {
	tl := TodoList{}

	id := uuid.New()
	todo := NewTodoWithID(id, "Test task", "This is a test")
	tl.AddTodo(todo)

	if len(tl.Todos) != 1 {
		t.Errorf("expected 1 todo in the todolist, got %d", len(tl.Todos))
	}

	if tl.Todos[0].ID != id {
		t.Errorf("expected ID %s, got %s", id, tl.Todos[0].ID)
	}
}

func TestRemoveTodo(t *testing.T) {
	tl := TodoList{}

	id := uuid.New()
	todo := NewTodoWithID(id, "Test task to remove", "Remove this task")
	tl.AddTodo(todo)

	if err := tl.RemoveTodo(id); err != nil {
		t.Fatalf("RemoveTodo returned an error: %v", err)
	}

	if len(tl.Todos) != 0 {
		t.Errorf("expected 0 todos in the todolist, got %d", len(tl.Todos))
	}
}

func TestUpdateTodoStatus(t *testing.T) {
	// We only test the status codes for now, not the strings
	expectedStatus := map[status]string{
		1: "doing",
		2: "done",
		0: "to do", // this is last because it's the initial status and we want to test the update
	}

	tl := TodoList{}

	id := uuid.New()
	todo := NewTodoWithID(id, "Test task", "Update my status")
	tl.AddTodo(todo)

	for wantStatusCode := range expectedStatus {
		if err := tl.UpdateTodoStatus(id, wantStatusCode); err != nil {
			t.Fatalf("UpdateTodoStatus returned an error: %v", err)
		}
		gotStatusCode := tl.Todos[0].Status // TODO: Make the test more robust with a GetTodo method ??
		if gotStatusCode != wantStatusCode {
			t.Errorf("expected %v status code, got %v", wantStatusCode, gotStatusCode)
		}
	}
}
