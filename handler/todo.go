package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hirac1220/go-clean-architecture/domain/model"
	"github.com/hirac1220/go-clean-architecture/usecase"
)

type TodoHandler interface {
	PostTodo(http.ResponseWriter, *http.Request)
	GetTodo(http.ResponseWriter, *http.Request)
	PutTodo(http.ResponseWriter, *http.Request)
	DeleteTodo(http.ResponseWriter, *http.Request)
	ListTodos(http.ResponseWriter, *http.Request)
}

type todoHandler struct {
	todoUseCase usecase.TodoUsecase
}

func NewTodoHandler(tu usecase.TodoUsecase) TodoHandler {
	return &todoHandler{
		todoUseCase: tu,
	}
}

func (th *todoHandler) PostTodo(w http.ResponseWriter, r *http.Request) {
	todo := &model.Todo{}
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	ctx := r.Context()
	t, err := th.todoUseCase.PostTodo(ctx, todo)
	if err != nil {
		if errors.Is(err, usecase.ErrNoAffected) {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	log.Printf("post todo: %v", t)

	res, err := json.Marshal(&t)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)

}

func (th *todoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("id: %v\n", vars["id"])

	ctx := r.Context()
	t, err := th.todoUseCase.GetTodo(ctx, vars["id"])
	if err != nil {
		if errors.Is(err, usecase.ErrNotFound) {
			log.Println(err)
			http.Error(w, "", http.StatusNotFound)
			return
		}
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	log.Printf("todo: %v", t)

	res, err := json.Marshal(&t)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (th *todoHandler) PutTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("id: %v\n", vars["id"])

	todo := &model.Todo{}
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	ctx := r.Context()
	a, err := th.todoUseCase.PutTodo(ctx, vars["id"], todo)
	if err != nil {
		if errors.Is(err, usecase.ErrNoAffected) {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	log.Printf("put todo: %v", todo)

	res, err := json.Marshal(&a)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (th *todoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("id: %v\n", vars["id"])

	ctx := r.Context()
	a, err := th.todoUseCase.DeleteTodo(ctx, vars["id"])
	if err != nil {
		if errors.Is(err, usecase.ErrNotFound) {
			log.Println(err)
			http.Error(w, "", http.StatusNotFound)
			return
		}
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	log.Printf("deleted affected: %v", a)

	res, err := json.Marshal(&a)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (th *todoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	t, err := th.todoUseCase.ListTodos(ctx)
	if err != nil {
		if errors.Is(err, usecase.ErrNotFound) {
			log.Println(err)
			http.Error(w, "", http.StatusNotFound)
			return
		}
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	log.Printf("todo list: %v", t)

	res, err := json.Marshal(&t)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)

}
