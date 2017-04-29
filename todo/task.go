package todo

import (
	"os"
	"log"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"github.com/satori/go.uuid"
)

type Task struct {
	Description string
}

func getDatastoreClient(ctx context.Context) *datastore.Client {
	projectId := os.Getenv("DATASTORE_PROJECT_ID")
	log.Printf(">>> DATASTORE_PROJECT_ID: %s", projectId)

	client, err := datastore.NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client
}

func CreateTask(desc string) Task {
	ctx := context.Background()
	client := getDatastoreClient(ctx)

	task := Task{
		Description: desc,
	}

	kind := "Task"
	name := uuid.NewV4().String()
	key  := datastore.NameKey(kind, name, nil)

	if _, err := client.Put(ctx, key, &task); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}

	return task
}
