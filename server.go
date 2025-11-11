package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const PORT  = ":8080"

type Todo struct {
	Id uint8 `json:"id"`
	Title string `json:"name"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {
	reader := Reader{filename: "test.json"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("reached 8080/")

		file, err := os.Open("test.json");
		if err != nil {
			log.Println("could not open test.json")
			file.Close()
			return
		}
		defer file.Close()

		fmt.Fprintln(w, "yo bro");
	})


	http.HandleFunc("/create", func(writer http.ResponseWriter, r *http.Request) {

		todo := Todo {
			Id: 0,
			Title: "my first todo",
			Content: "un contenu comme un autre",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		reader.Create(todo)

		fmt.Fprintln(writer, "Todo created");
	
	})



	fmt.Println("starting server listening on port",PORT)
	http.ListenAndServe(PORT, nil)
}
