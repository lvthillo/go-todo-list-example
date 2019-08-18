package todo

import (
	"errors"
	"sync"

	"github.com/rs/xid"
)

var (
	// list which will contain todos
	list []Todo
	// this is the mutex that will allow you to safely access/manipulate the data in this package across different goroutines.
	// goroutines are multiple lightweight threads of execution.
	mtx sync.RWMutex
	// assure that a specific operation will run only once.
	once sync.Once
)

// the init function will run when the package is initialized.
// it will trigger another function 'initialiseList'.
// by using the 'once.Do' we can assure that the list will only be initialized once and not during every initialization of the package.
func init() {
	once.Do(initialiseList)
}

// triggered by the init function.
// this will just return an empty Todo list (slice)
func initialiseList() {
	list = []Todo{}
}

// Todo data structure for a task with a description of what to do
// every Todo struct will contain an ID, a Message and a Complete status.
// the second part of the struct shows how these 'fields' can be accessed in JSON.
// Todo with capital so we can access this struct in another module.
type Todo struct {
	ID       string `json:"id"`
	Message  string `json:"message"`
	Complete bool   `json:"complete"`
}

// Get retrieves all elements from the todo list
// function is public and will return the current todo list.
func Get() []Todo {
	return list
}

// Add will add a new todo based on a message
// use the mutex here to lock and unlock your message creation to avoid a race condition.
// as your server might handle multiple operations at the same time and these operations try to access the same memory, you can run into a race-condition that might make Golang crash.
// public function.
func Add(message string) string {
	// use newTodo function to create a new Todo struct.
	t := newTodo(message)
	mtx.Lock()
	// append the Todo item to the Todo list.
	list = append(list, t)
	mtx.Unlock()
	// return the ID of the new struct item.
	return t.ID
}

// Delete will remove a Todo from the Todo list based on ID.
// public function.
func Delete(id string) error {
	// use findTodoLocation to find the location of the ToDo item in the list.
	location, err := findTodoLocation(id)
	if err != nil {
		return err
	}
	// remove element when it exists using removeElementByLocation function.
	removeElementByLocation(location)
	return nil
}

// Complete will set the complete boolean to true, marking a todo as completed.
// comparable with Delete function.
// public function.
func Complete(id string) error {
	// same as Delete, find Todo in list based ID.
	location, err := findTodoLocation(id)
	if err != nil {
		return err
	}
	// Set todo complete using the setTodoCompleteByLocation function.
	setTodoCompleteByLocation(location)
	return nil
}

// create a new Todo item following the struct
// an ID will be set and a message (parameter)
func newTodo(msg string) Todo {
	return Todo{
		ID:       xid.New().String(),
		Message:  msg,
		Complete: false,
	}
}

// function which returns the location of a Todo struct in the list of Todo's.
func findTodoLocation(id string) (int, error) {
	// Lock for reading (we are only reading, not writing here).
	mtx.RLock()
	// Unlock
	// Defer: A defer statement defers the execution of a function until the surrounding function returns.
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	defer mtx.RUnlock()
	// loop through our list (slice) and check when the t.ID (= ID field of struct (t)) matches the id parameter using the isMatchingID function.
	for i, t := range list {
		if isMatchingID(t.ID, id) {
			return i, nil
		}
	}
	return 0, errors.New("could not find todo based on id")
}

// remove Todo struct based on location in the list (= i parameter)
func removeElementByLocation(i int) {
	// lock
	mtx.Lock()
	// Set the list to a new array, which contains all elements from the previous list up to a given location + appended with all elements after (but not including) the same given location.
	// We will get a new list without that given location (essentially deleting it from the previous list).
	list = append(list[:i], list[i+1:]...)
	// unlock
	mtx.Unlock()
}

// setTodoCompleteByLocation sets .Complete variable of struct (based on location) in list on true.
func setTodoCompleteByLocation(location int) {
	// lock
	mtx.Lock()
	// find struct in list based by loction and set it on true.
	list[location].Complete = true
	// unlock
	mtx.Unlock()
}

// isMatchingID compares two strings (ID's) and when they match it will return true.
func isMatchingID(a string, b string) bool {
	return a == b
}
