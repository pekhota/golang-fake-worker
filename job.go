package main

import "github.com/google/uuid"

type Job struct {
	Id string `json:"id"`
	Status string `json:"status"`
	Payload string `json:"payload"`
	CallbackUrl string `json:"callback_url"`
}

const JobStatusWaiting = "waiting"
const JobStatusInProgress  = "progress"
const JobStatusDone  = "done"

var jobs []Job

func initJobs()  {
	jobs = []Job{
		{uuid.New().String(), JobStatusDone, "hello world", ""},
		{uuid.New().String(), JobStatusDone, "hello world2", ""},
	}

}
