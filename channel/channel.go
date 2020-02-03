package channel

import (
	"sync"

	model "github.com/alterra/graphql-server/models"
)

var once sync.Once

var (
	instance chan *model.Todo
)

func NewClass() chan *model.Todo {

	once.Do(func() { // <-- atomic, does not allow repeating

		instance = make(chan *model.Todo, 100) // <-- thread safe

	})

	return instance
}
