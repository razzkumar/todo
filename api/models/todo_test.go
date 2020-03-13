package models

import "testing"

func TestAdd(t *testing.T) {
	todo := New()

	todo.Add(Todo{ID: 1, Data: "test todo", IsDone: false, IsStared: true})

	if len(todo.Todos) != 1 {
		t.Errorf("Todo Was not added")
	}
}
func TestGetAll(t *testing.T) {
	todo := New()
	todo.Add(Todo{})
	result := todo.GetAll()

	if len(result) != 1 {
		t.Errorf("Todo was not added")
	}
}
