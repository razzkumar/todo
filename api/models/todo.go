package models

import "time"

type Getter interface {
	GetAll() []Todo
}

type Adder interface {
	Add(todo Todo)
}

type Todo struct {
	ID        int       `json:"id"`
	Data      string    `json:"data"`
	IsDone    bool      `json:"isDone"`
	CreatedAt time.Time `json:"create_at"`
	IsStared  bool      `json:"isStared"`
}

type TodoList struct {
	Todos []Todo
}

func New() *TodoList {
	return &TodoList{
		Todos: []Todo{},
	}
}

func (r *TodoList) Add(todo Todo) {
	r.Todos = append(r.Todos, todo)
}

func (r *TodoList) GetAll() []Todo {
	return r.Todos
}
