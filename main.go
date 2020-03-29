package main

import (
	"log"
	"net/http"
)

func handleError(err error) {
	if err != nil {
		log.Println("handle error")
		panic(err)
	}
}

func main()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	go callbackSender()

	router := NewRouter()

	log.Fatal("Fatal", http.ListenAndServe(":10000", router))
}