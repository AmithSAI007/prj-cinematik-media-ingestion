package metadata

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AmithSAI007/prj-cinematik-pubsub-metadata.git/internal/pubsub"
	publisher "github.com/AmithSAI007/prj-cinematik-pubsub-metadata.git/pkg/pubsub"
	"github.com/AmithSAI007/prj-cinematik-pubsub-metadata.git/pkg/transform"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

// Environment variable keys
const (
	envProjectID = "PROJECT_ID"
	envTopicID   = "TOPIC_ID"
)

// init registers the CloudEvent function handler.
func init() {
	functions.CloudEvent("PubSubMetadata", pubsubMetadata)
}

// pubsubMetadata handles CloudEvents by processing storage object metadata
// and publishing it to a specified Pub/Sub topic.
//
// Parameters:
//   - ctx: The context.Context for the function execution
//   - e: The CloudEvent containing the storage object data
//
// Returns:
//   - error: nil if successful, error otherwise
//
// The function expects the following environment variables to be set:
//   - PROJECT_ID: The Google Cloud Project ID
//   - TOPIC_ID: The Pub/Sub topic ID to publish messages to
func pubsubMetadata(ctx context.Context, e event.Event) error {
	log.Printf("Event ID: %s", e.ID())
	log.Printf("Event Type: %s", e.Type())

	var data transform.StorageObjectData
	projectId := os.Getenv(envProjectID)
	topicId := os.Getenv(envTopicID)

	if projectId == "" || topicId == "" {
		return fmt.Errorf("required environment variables not set: PROJECT_ID and/or TOPIC_ID")
	}

	if err := e.DataAs(&data); err != nil {
		return fmt.Errorf("unable to process event data: %v", err)
	}

	log.Printf("Inititalizing Pub/Sub client, project id: %s", projectId)
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	log.Printf("Initializing publisher and preparing message")
	publisher := publisher.NewPublisher(client.Client)

	msg := transform.TransformToTopicMessageData(data)

	log.Printf("Attempting to publish message to topic: %s", topicId)
	id, err := publisher.Publish(ctx, topicId, msg)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("published message with id: %s", id)
	return nil
}
