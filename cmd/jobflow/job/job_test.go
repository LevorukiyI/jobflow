package job

import (
	"bytes"
	"testing"
)

/*
create
get
not found
delete
delete missing
list
find by type
start
complete
fail
validation
invalid state transition
*/

func TestNewJob(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		jobType string
		payload []byte
		check   func(*testing.T, *Job)
	}{
		{
			name:    "test id creation",
			id:      "job-1",
			jobType: "jobtype-1",
			payload: []byte{},
			check: func(t *testing.T, j *Job) {
				var want string = "job-1"
				if want != j.ID {
					t.Errorf("ID = %q, want %q", j.ID, want)
				}
			},
		},
		{
			name:    "test type createion",
			id:      "job-1",
			jobType: "jobtype-1",
			payload: []byte{},
			check: func(t *testing.T, j *Job) {
				var want string = "jobtype-1"
				if want != j.Type {
					t.Errorf("Type = %q, want %q", j.Type, want)
				}
			},
		},
		{
			name:    "test payload creation",
			id:      "job-1",
			jobType: "jobtype-1",
			payload: []byte("some string"),
			check: func(t *testing.T, j *Job) {
				var want []byte = []byte("some string")
				if !bytes.Equal(j.Payload, want) {
					t.Errorf("Payload = %q, want = %q", j.Payload, want)
				}
			},
		},
		{
			name:    "test status pending after creation",
			id:      "job-1",
			jobType: "jobtype-1",
			payload: []byte("some string"),
			check: func(t *testing.T, j *Job) {
				if j.Status != StatusPending {
					t.Errorf("Status = %q, want = %q", j.Status, StatusPending)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			job := NewJob(tt.id, tt.jobType, tt.payload)

			tt.check(t, &job)
		})
	}
}

func TestJobComplete(t *testing.T) {
	job := NewJob("job-1", "email", nil)

	job.Complete()

	if job.Status != StatusCompleted {
		t.Fatalf(
			"expected %q, got %q",
			StatusCompleted,
			job.Status,
		)
	}
}
