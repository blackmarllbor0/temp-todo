package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// GenerateId генерирует новый id
func GenerateId() string {
	b := make([]byte, 16)

	if _, err := rand.Read(b); err != nil {
		log.Fatal(err.Error())
	}

	return fmt.Sprintf("%x", b)
}

// generateTemplate создает шаблон с переданными данными
func generateTemplate(w http.ResponseWriter, filename, define string, data interface{}) {
	temp, err := template.ParseFiles(filename, header, footer)
	if err != nil {
		log.Fatal(err)
	}
	if err := temp.ExecuteTemplate(w, define, data); err != nil {
		log.Fatal(err)
	}
}
