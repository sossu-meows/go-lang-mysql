package main

import (
	"encoding/json"
	"first-app/package/lib"
	"log"
	"net/http"
)

func main() {
	// capture connection

	lib.ConnectToMysql()

	// fmt.Println("Hello World!")
	// testFunc()
	// test := quote.Go()
	// fmt.Println(test)
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ngay day"))
	})

	http.HandleFunc("/getBook", func(w http.ResponseWriter, r *http.Request) {
		GetBookJson(w, r)
	})

	http.HandleFunc("/getBooks", GetBookJson)

	log.Fatal(
		http.ListenAndServe(":5000", nil))
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func GetBookJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	book := Book{
		Title:  "The hobbit",
		Author: "Khoa",
		Pages:  156,
	}

	json.NewEncoder(w).Encode(book)
}
