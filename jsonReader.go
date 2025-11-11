package main

import (
	"encoding/json"
	"log"
	"os"
)

type JsonReader interface {
	openFile() *os.File
	GetAll() []Todo
	get(id uint8)
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

func (reader Reader) GetAll() []Todo {
	file := reader.openFile()
	defer file.Close()
	var todos []Todo
	err := json.NewDecoder(file).Decode(&todos)

	if err != nil && err.Error() != "EOF" {
		log.Fatal("couldnt decode the file", err);
	}

	println(todos)
	return todos
}

func (reader Reader) Create(todo Todo) {
	file := reader.openFile()
	defer file.Close()

	todos := reader.GetAll()

	todos = append(todos, todo)
	
	_, err := file.Seek(0, 0)
	println("bruhg")

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

func NewReader() *Reader {
	return &Reader{
		filename: "test.json",
	}
}
