package gcp

import (
	"go_pubsub_job/internal/infrastructure/ctx"

	"cloud.google.com/go/firestore"
)

type FirestoreMapping interface {
	Doc() string

	ToMap() map[string]any
}

func FirestoreSave(client *firestore.Client, collectionName string, data FirestoreMapping) error {
	_ctx, cancel := ctx.NewTimeoutContext()
	defer cancel()

	_, err := client.Collection(collectionName).
		Doc(data.Doc()).
		Set(_ctx, data.ToMap(), firestore.MergeAll)
	if err != nil {
		return err
	}

	return nil
}
