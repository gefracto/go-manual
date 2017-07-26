package main

import (
	"net/http"
	"os"
)

func main() {
	fileserver := http.FileServer(http.Dir("static"))
	http.Handle("/", fileserver)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
