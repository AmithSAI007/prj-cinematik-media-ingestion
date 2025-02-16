package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/AmithSAI007/prj-cinematik-pubsub-metadata.git/pkg/transform"
	"go.uber.org/zap"
)

// Publisher is responsible for publishing messages to Google Cloud Pub/Sub topics.
type Publisher struct {
	client *pubsub.Client
	log    *zap.Logger
}

// NewPublisher creates a new Publisher instance using an existing Pub/Sub client.
func NewPublisher(client *pubsub.Client, log *zap.Logger) *Publisher {
	return &Publisher{
		client: client,
		log:    log,
	}
}

// validateTopic checks whether the specified Pub/Sub topic exists.
// It returns an error if the topic does not exist or if an error occurs while checking.
func (p *Publisher) validateTopic(ctx context.Context, topicId string) error {
	topic := p.client.Topic(topicId)
	exists, err := topic.Exists(ctx)
	if err != nil {
		return fmt.Errorf("failed to check if topic exists: %w", err)
	}

	if !exists {
		return fmt.Errorf("topic does not exists: %s", topicId)
	}

	return nil
}

// createMessage converts a TopicMessgaeData object into a Pub/Sub message.
// It returns the created message or an error if marshalling fails.
func (p *Publisher) createMessage(msg transform.TopicMessgaeData) (*pubsub.Message, error) {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal message: %w", err)
	}

	return &pubsub.Message{
		Data: msgBytes,
		Attributes: map[string]string{
			"origin":   "golang",
			"username": "cinematik",
		},
	}, nil
}

// Publish publishes a message to the specified Pub/Sub topic.
// It validates the topic, creates the message, and then sends it to the topic.
// It returns the message ID if successful or an error if any step fails.
func (p *Publisher) Publish(ctx context.Context, topicId string, msg transform.TopicMessgaeData) (string, error) {
	p.log.Info("Attempting to publish message",
		zap.String("topicId", topicId),
		zap.String("bucket", msg.Bucket),
		zap.String("file", msg.FileName),
	)
	if err := p.validateTopic(ctx, topicId); err != nil {
		return "", err
	}

	pubsubMsg, err := p.createMessage(msg)
	if err != nil {
		return "", err
	}

	topic := p.client.Topic(topicId)
	result := topic.Publish(ctx, pubsubMsg)

	id, err := result.Get(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to publish message: %w", err)
	}

	return id, nil
}
