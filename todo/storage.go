package todo

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

func checkLogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitTodosFile() {
	if _, err := os.Stat("todos.json"); !errors.Is(err, os.ErrNotExist) {
		checkLogError(err)
		return
	}

	f, err := os.Create("todos.json")
	checkLogError(err)

	defer f.Close()

	if _, err := f.WriteString("[]\n"); err != nil {
		log.Fatal(err)
	}
}

func ReadTodosFile() []Todo {
	b, err := os.ReadFile("todos.json")
	checkLogError(err)

	var todos []Todo

	if err := json.Unmarshal(b, &todos); err != nil {
		log.Fatal(err)
	}

	return todos
}
