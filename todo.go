package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) {

	newTodo := Todo{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	*todos = append(*todos, newTodo)
}

// func (todos Todos) Add(title string) {

// 	newTodo := Todo{
// 		title:     title,
// 		completed: false,
// 		createdAt: time.Now(),
// 	}

// 	todos = append(todos, newTodo)
// } // when using value receiver, the changes won't reflect outside, because a copy is made, that is usable only inside the method.

/*
*
*
Key Reasons:
Modify the original: With *Todos, any changes to the slice are made to the original data, not a copy
Memory efficiency: Avoids copying the entire slice when the method is called
Consistency: Since your Add method uses *Todos and modifies the slice, validateIndex should also use *Todos for consistency

// With pointer receiver (*Todos)

	func (todos *Todos) Add(title string) {
	    *todos = append(*todos, newTodo)  // Modifies original
	}

// If you used value receiver (Todos)

	func (todos Todos) Add(title string) {
	    todos = append(todos, newTodo)  // Only modifies the copy, original unchanged!
	}

For validateIndex:
While validateIndex doesn't modify the slice, using *Todos is still good practice because:

Consistency: All methods on the same type should use the same receiver type
Performance: Avoids copying the slice, especially if it becomes large
General Rule: Use pointer receivers when you need to modify the receiver OR when the receiver is a large struct/slice.

*
*/
func (todos *Todos) validateIndex(index int) error {

	if index < 0 || index >= (len(*todos)) {
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) delete(index int) error {

	todo := *todos

	// err := todo.validateIndex(index)
	// if err != nil {
	// 	return err
	// }

	if err := todo.validateIndex(index); err != nil {
		return err
	}

	*todos = append(todo[:index], todo[index+1:]...)

	return nil
}

func (todos *Todos) toggle(index int) error {

	todo := *todos

	err := todo.validateIndex(index)

	if err != nil {
		return err
	}

	isCompleted := todo[index].Completed

	if !isCompleted {
		completedAt := time.Now()
		todo[index].CompletedAt = &completedAt

	}

	todo[index].Completed = !isCompleted
	return nil

}

// without dereferencing, indexing dirctly on pointers --------------------------------------------

// func (todos *Todos) toggle1(index int) error {
// 	if err := todos.validateIndex(index); err != nil {
// 		return err
// 	}

// 	// Index directly on the pointer
// 	isCompleted := (*todos)[index].Completed

// 	if !isCompleted {
// 		now := time.Now()
// 		(*todos)[index].CompletedAt = &now
// 	}

// 	(*todos)[index].Completed = !isCompleted

// 	return nil
// }

func (todos *Todos) edit(index int, title string) error {

	todoSlice := *todos

	err := todoSlice.validateIndex(index)

	if err != nil {
		return err
	}

	todoSlice[index].Title = title
	return nil
}

// now install this Go get github.com/aquasecurity/table  for visually appealing table.

func (todos *Todos) print() {

	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}
