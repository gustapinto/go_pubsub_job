package job

import (
	"go_pubsub_job/internal/domain/job"
	"go_pubsub_job/internal/infrastructure/gcp"

	"cloud.google.com/go/firestore"
)

type JobStateFirestoreRepository struct {
	Client *firestore.Client
}

func (r *JobStateFirestoreRepository) Save(state job.JobState) error {
	return gcp.FirestoreSave(r.Client, "job_states", &state)
}
