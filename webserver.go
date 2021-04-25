package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hallo!")

	fmt.Println("Der Server wird erstellt...")

	http.Handle("/", http.FileServer(http.Dir("static")))

	http.ListenAndServe(":8080", nil)

	fmt.Println("Server fertig. Bitte öffne im Browser: localhost:8080 oder http://127.0.0.1:8080")
}
