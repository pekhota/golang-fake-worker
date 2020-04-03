package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)


func handleError(err error) {
	if err != nil {
		log.Println("handle error")
		panic(err)
	}
}

func main()  {
	port := flag.Int("port", 9292, "Default part to run server on")

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.Parse()


	go callbackSender()

	initJobs()

	router := NewRouter()


	addr := ":" + strconv.Itoa(*port)
	fmt.Println("Listening at " + addr)
	log.Fatal("Fatal", http.ListenAndServe(addr, router))
}