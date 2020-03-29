package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"net/http"
)

func makeJob(w http.ResponseWriter, r *http.Request)  {

	reqBody, err := ioutil.ReadAll(io.LimitReader(r.Body, 1000*1000))

	var job Job

	if err = json.Unmarshal(reqBody, &job); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		err := json.NewEncoder(w).Encode(err);
		handleError(err)
		return
	}

	job.Status = JobStatusWaiting
	job.Id = uuid.New().String();

	jobs = append(jobs, job)

	index := len(jobs) - 1
	go worker(index);

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(job)
}

func homeHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello from async worker")
}

func listJobs(w http.ResponseWriter, r *http.Request)  {
	json.NewEncoder(w).Encode(jobs)
}
