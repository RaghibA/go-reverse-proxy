package proxy

import (
	"log"
	"net/http"
)

func Run() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Proxy server running on port 8080")
}
