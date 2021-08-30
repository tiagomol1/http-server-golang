package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/tiagomol1/http-server-golang/router"
)

func main() {
	router.Router()

	http.Handle("/", http.FileServer(http.Dir("./public")))

	fmt.Println("> Server running on port 8081")
	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
