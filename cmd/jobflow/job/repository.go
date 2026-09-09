package job

import "fmt"

type JobRepository interface {
	Add(Job) error
	Get(Job) (Job, error)
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
