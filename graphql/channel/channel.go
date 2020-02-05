package channel

import (
	"sync"

	model "github.com/alterra/graphql-server/graphql/models"
)

var once sync.Once

var (
	channels map[string]chan *model.Todo
	mutex    sync.Mutex
)

func NewChannel(key string) chan *model.Todo {
	if channels == nil {
		channels = make(map[string]chan *model.Todo)
	}

	channel := make(chan *model.Todo, 1)
	mutex.Lock()
	channels[key] = channel
	mutex.Unlock()

	return channel
}

func DeleteChannel(key string) {
	mutex.Lock()
	delete(channels, key)
	mutex.Unlock()
}

func PublishAll(todo *model.Todo) {
	for _, v := range channels {
		v <- todo
	}
}
