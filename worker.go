package main

import (
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

var c chan int = make(chan int)

func callbackSender()  {
	for  {
		index := <- c

		go func() {
			_, err := http.PostForm(jobs[index].CallbackUrl, url.Values{"key": {"Value"}, "id": {"123"}})
			if err != nil {
				log.Println("Line 50", err)
			}
		}()
	}

}


func worker(jobIndex int)  {
	// todo for demo make delay in range of 60-90 seconds
	jobs[jobIndex].Status = JobStatusInProgress

	amt := time.Duration(rand.Intn(10)+1)
	time.Sleep(time.Second * amt)

	jobs[jobIndex].Status = JobStatusDone

	c <- jobIndex
}
