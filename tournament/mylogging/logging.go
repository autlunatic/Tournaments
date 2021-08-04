package mylogging

import (
	"context"
	"log"

	"cloud.google.com/go/logging"
)

func AddLog(s ...interface{}) {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "kids-192720"

	// Creates a client.
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Sets the name of the log to write to.
	logName := "my-log"

	logger := client.Logger(logName).StandardLogger(logging.Info)

	// Logs "hello world", log entry is visible at
	// Cloud Logs.
	logger.Println(s...)
}
