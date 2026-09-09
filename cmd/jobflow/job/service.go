package job

type JobService struct {
	repository JobRepository
}

func NewJobService(repository JobRepository) *JobService {
	return &JobService{
		repository: repository,
	}
}
