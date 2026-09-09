package job

import (
	"errors"
	"time"
)

type Job struct {
	ID        string
	Type      string
	Payload   []byte
	Status    JobStatus
	CreatedAt time.Time
}

type JobStatus string

const (
	StatusPending   JobStatus = "pending"
	StatusRunning   JobStatus = "running"
	StatusCompleted JobStatus = "completed"
	StatusFailed    JobStatus = "failed"
)

func NewJob(
	id string,
	jobType string,
	payload []byte,
) Job {
	return Job{
		ID:        id,
		Type:      jobType,
		Payload:   payload,
		Status:    StatusPending,
		CreatedAt: time.Now(),
	}
}

func SetStatus(status JobStatus) {

}

var ErrJobNotFound = errors.New("job not found")

func FindJob(
	jobs []Job,
	id string,
) (job *Job) {
	for i := range jobs {
		if jobs[i].ID == id {
			job = &jobs[i]
			return job
		}
	}
	return nil
}

func (j *Job) IsFinished() bool {
	return j.Status == StatusCompleted ||
		j.Status == StatusFailed
}

func (j *Job) Complete() {
	j.Status = StatusCompleted
}
