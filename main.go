package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/hirac1220/go-clean-architecture/domain/repository"
	"github.com/hirac1220/go-clean-architecture/handler"
	"github.com/hirac1220/go-clean-architecture/infrastructure/persistence"
	"github.com/hirac1220/go-clean-architecture/usecase"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/cors"
	"github.com/tinrab/retry"
)

type Config struct {
}

var tu usecase.TodoUsecase

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var tp repository.TodoRepository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		tp, err = persistence.NewTodoPersistence()
		if err != nil {
			log.Println(err)
		}
		return err
	})
	defer tp.Close()

	tu = usecase.NewTodoUseCase(tp)
	th := handler.NewTodoHandler(tu)

	// Routing
	r := mux.NewRouter()
	r.Use(Middleware)

	r.HandleFunc("/todos/{userId}", th.PostTodo).Methods("POST")
	r.HandleFunc("/todos/{userId}/{id}", th.GetTodo).Methods("GET")
	r.HandleFunc("/todos/{userId}/{id}", th.PutTodo).Methods("PUT")
	r.HandleFunc("/todos/{userId}/{id}", th.DeleteTodo).Methods("DELETE")
	r.HandleFunc("/todos/{userId}", th.ListTodos).Methods("GET")

	// Set cors
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	})

	log.Fatal(http.ListenAndServe(":8080", c.Handler(r)))
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("user_id: %v\n", vars["userId"])
		ctx := r.Context()
		id, err := tu.CheckUser(ctx, vars["userId"])
		if err == nil {
			log.Println("user_id:%d is match", id)
			next.ServeHTTP(w, r)
		} else {
			if errors.Is(err, usecase.ErrNotFound) {
				log.Println(err)
				http.Error(w, "user is not exist", http.StatusNotFound)
				return
			}
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	})
}
