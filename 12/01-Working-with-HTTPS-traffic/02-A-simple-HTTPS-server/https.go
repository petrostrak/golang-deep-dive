package main

import (
	"fmt"
	"net/http"
)

var (
	PORT = ":8080"
)

func Default(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is an example HTTPS server!\n")
}

func main() {

	http.HandleFunc("/", Default)
	fmt.Println("Listening to port number", PORT)

	if err := http.ListenAndServeTLS(PORT, "server.crt", "server.key", nil); err != nil {
		fmt.Println(err)
		return
	}

}
