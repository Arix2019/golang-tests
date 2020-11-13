package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Wellcome!")
	log.Printf("handler started.")
	defer log.Printf("handler ended.")

	time.Sleep(3 * time.Second) // aguarda 3 seg. e finaliza a conexão
	fmt.Fprintln(w, "Bye!")
}
