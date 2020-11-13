package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() //fecha o corpo da resposta
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	io.Copy(os.Stdout, res.Body) //copia para o suporte de saida o corpo da resposta
}
