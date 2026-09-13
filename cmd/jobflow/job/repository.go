package job

import "fmt"

type JobRepository interface {
	Add(Job) error
	Get(Job) (Job, error)
	Delete(string) error
	List() []Job
}

type MemoryJobRepository struct {
	jobs map[string]Job
}

func NewMemoryJobRepository() *MemoryJobRepository {
	return &MemoryJobRepository{
		jobs: make(map[string]Job),
	}
}

func (repository *MemoryJobRepository) Get(
	id string,
) (Job, error) {
	job, ok := repository.jobs[id]
	if !ok {
		return Job{}, fmt.Errorf(
			"get job %s: %w",
			id,
			ErrJobNotFound,
		)
	}
	return job, nil
}

func (repository *MemoryJobRepository) Add(job Job) error {
	repository.jobs[job.ID] = job
	return nil
}
