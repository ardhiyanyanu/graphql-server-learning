package service

import (
	"strconv"

	model "github.com/alterra/graphql-server/models"
)

var temp []*model.Todo
var count = 0

func CreateTodo(input model.NewTodo) (*model.Todo, error) {
	todo := model.Todo{
		ID:   "x" + strconv.FormatInt(int64(count), 10),
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID: input.UserID,
		},
	}
	temp = append(temp, &todo)
	count = count + 1
	return &todo, nil
}

func Todos() ([]*model.Todo, error) {
	return temp, nil
}
