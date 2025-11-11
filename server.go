package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const PORT = ":8080"

type Todo struct {
	Id        uint8     `json:"id"`
	Title     string    `json:"name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {
	reader := Reader{filename: "test.json"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("reached 8080/")

		file, err := os.Open("test.json")
		if err != nil {
			log.Println("could not open test.json")
			file.Close()
			return
		}
		defer file.Close()

		fmt.Fprintln(w, "yo bro")
	})

	http.HandleFunc("/create", func(writer http.ResponseWriter, r *http.Request) {
		todo := Todo{
			Title:     "my first todo",
			Content:   "un contenu comme un autre",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		reader.Create(todo)
		fmt.Fprintln(writer, "Todo created")
	})

	http.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		param := r.PathValue("id")

		if param == "" {
			log.Println("id parameter is missing")
			http.Error(w, "id parameter is missing", http.StatusBadRequest)
			return
		}

		println("param:", param)
		id, err := strconv.Atoi(param)

		if err != nil {
			log.Println("invalid id parameter", err)
			log.Fatal("invalid id parameter", err)
			http.Error(w, "invalid id parameter", http.StatusBadRequest)
			return
		}

		todo, success := reader.Get(uint8(id))

		if !success {
			log.Println("todo not found")
			http.Error(w, "todo not found", http.StatusNotFound)
			return
		}

		log.Println("reached 8080/{id}")
		fmt.Fprintln(w, "Todo:", todo)
	})

	http.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
		todos, err := reader.GetAll()
		if err != nil {
			log.Println("could not get all todos:", err)
			http.Error(w, "could not get all todos", http.StatusInternalServerError)
			return 
		}

		log.Println("reached 8080/all")
		fmt.Fprintln(w, "Todos:", todos)
	})

	fmt.Println("starting server listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}
