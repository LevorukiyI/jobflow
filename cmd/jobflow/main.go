package main

import (
	"fmt"
	"jobflow/cmd/jobflow/job"
)

func main() {
	fmt.Println("test message for go code +==============================")
	var jobs []job.Job = []job.Job{
		job.NewJob("job1", "jobType", []byte{}),
		job.NewJob("job2", "jobType", []byte{}),
		job.NewJob("job3", "jobType", []byte{}),
	}
	for _, job := range jobs {
		fmt.Println(job.ID)
	}
	foundedJob := job.FindJob(jobs, "job2")
	fmt.Println(foundedJob.ID)

}
