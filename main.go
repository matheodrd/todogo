package main

import (
	"fmt"

	"github.com/matheodrd/todogo/cmd"
	"github.com/matheodrd/todogo/todo"
)

func main() {
	todo.InitTodosFile()
	todos := todo.ReadTodosFile()
	fmt.Println(todos[0].Status)
	cmd.Execute()
}
