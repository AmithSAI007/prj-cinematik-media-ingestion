package metadata

import (
	"context"
	"fmt"
	"os"

	"github.com/AmithSAI007/prj-cinematik-pubsub-metadata.git/internal/pubsub"
	publisher "github.com/AmithSAI007/prj-cinematik-pubsub-metadata.git/pkg/pubsub"
	"github.com/AmithSAI007/prj-cinematik-pubsub-metadata.git/pkg/transform"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"go.uber.org/zap"
)

// Environment variable keys
const (
	envProjectID = "GCP_PROJECT_ID"
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
// Tme function expects the following environment variables to be set:
//   - PROJECT_ID: The Google Cloud Project ID
//   - TOPIC_ID: The Pub/Sub topic ID to publish messages to
func pubsubMetadata(ctx context.Context, e event.Event) error {
	log, err := zap.NewProduction()
	if err != nil {
		return fmt.Errorf("unable to initialize zap logging, %v", err)
	}
	defer log.Sync()

	log.Info("Received CloutEvent",
		zap.String("EventID", e.ID()),
		zap.String("EventType", e.Type()),
	)

	var data transform.StorageObjectData
	projectId := os.Getenv(envProjectID)
	topicId := os.Getenv(envTopicID)

	if projectId == "" || topicId == "" {
		return fmt.Errorf("required environment variables not set: PROJECT_ID and/or TOPIC_ID")
	}

	if err := e.DataAs(&data); err != nil {
		return fmt.Errorf("unable to process event data: %v", err)
	}

	log.Info("Inititalizing Pub/Sub client",
		zap.String("Project ID", projectId),
	)
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		return fmt.Errorf("failed to create pub/sub client: %v", err)
	}
	defer client.Close()

	publisher := publisher.NewPublisher(client.Client, log)

	msg, err := transform.TransformToTopicMessageData(data)
	if err != nil {
		return fmt.Errorf("failed to transform message: %v", err)
	}

	log.Info("Attempting to publish message",
		zap.String("topicId", topicId),
	)
	id, err := publisher.Publish(ctx, topicId, msg)
	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	log.Info("Message published successfully",
		zap.String("messageId", id),
	)
	return nil
}
