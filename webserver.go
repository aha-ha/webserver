package main

import (
	"fmt"
	"net/http"
)

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Erstellt mit Aha-Ha Webservercreator <a href='https://aha-ha.github.io' target='_blank'>Website</a>")
}

func main() {
	fmt.Println("Hallo!")

	fmt.Println("Der Server wird erstellt...")

	http.HandleFunc("/about", about)

	http.Handle("/", http.FileServer(http.Dir("static")))

	http.ListenAndServe(":8080", nil)

	fmt.Println("Server fertig. Bitte öffne im Browser: localhost:8080 oder http://127.0.0.1:8080")
}
