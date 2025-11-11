package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type JsonReader interface {
	openFile() *os.File
	GetAll() []Todo
	Get(id uint8) Todo
	update(id uint8)
	remove(id uint)
	Create() Todo
}

type Reader struct {
	filename string
}

func (reader Reader) openFile() *os.File {
	file, err := os.OpenFile(reader.filename, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal("could not open file: ", err)
	}

	return file
}

func (reader Reader) GetAll() ([]Todo, error) {
	file := reader.openFile()
	defer file.Close()
	var todos []Todo
	err := json.NewDecoder(file).Decode(&todos)

	if err != nil && err.Error() != "EOF" {
		return nil, err
	}

	return todos, nil
}

func (reader Reader) Create(todo Todo) {
	file := reader.openFile()
	defer file.Close()

	todos, err := reader.GetAll()

	if err != nil {
		log.Fatal("could not get all todos: ", err)
	}

	var id uint8
	if len(todos) == 0 {
		id = 1
	} else {
		id = todos[len(todos)-1].Id + 1
	}

	todo.Id = id

	todos = append(todos, todo)

	_, err = file.Seek(0, 0)

	if err != nil && err.Error() != "EOF" {
		log.Fatal("ici", err)
	}

	err = file.Truncate(0)

	if err != nil {
		log.Fatal("could not replace the content of the file :", err)
	}

	err = json.NewEncoder(file).Encode(&todos)

	if err != nil {
		log.Fatal("could not write to file: ", err)
	}
}

func (reader Reader) Get(id uint8) (Todo, bool) {
	file := reader.openFile()
	defer file.Close()
	todos, err := reader.GetAll()

	if err != nil {
		log.Fatal("could not get all todos: ", err)
	}

	for _, todo := range todos {
		if todo.Id == id {
			return todo, true
		}
	}
	return Todo{}, false
}

func (reader Reader) update(id uint8) {
	file := reader.openFile()
	defer file.Close()

	var todos []Todo
	err := json.NewDecoder(file).Decode(&todos)
	if err != nil && err.Error() != "EOF" {
		log.Fatal("couldnt decode the file", err)
	}

	for i, todo := range todos {
		if todo.Id == id {
			todos[i].UpdatedAt = time.Now()
		}
	}
}

func NewReader() *Reader {
	return &Reader{
		filename: "test.json",
	}
}
