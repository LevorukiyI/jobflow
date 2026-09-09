package job

import (
	"errors"
	"fmt"
)

var ErrInvalidJob = errors.New("invalid job")

func ValidateJob(job Job) error {
	if job.ID == "" {
		return fmt.Errorf("job Id: %w", ErrInvalidJob)
	}
	if job.Type == "" {
		return fmt.Errorf("job Type: %w", ErrInvalidJob)
	}

	return nil
}
