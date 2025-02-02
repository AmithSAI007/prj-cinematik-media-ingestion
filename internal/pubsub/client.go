package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

// Client represents a Pub/Sub client, holding the project ID
// and the actual Google Cloud Pub/Sub client.
type Client struct {
	ProjectId string
	Client    *pubsub.Client
}

// NewClient creates and returns a new Client for interacting with Google Cloud Pub/Sub.
// It accepts a context and a project ID to initialize the Pub/Sub client.
// Returns a pointer to a Client instance and an error if any occurs during initialization.
func NewClient(ctx context.Context, projectId string) (*Client, error) {
	client, err := pubsub.NewClientWithConfig(ctx, projectId, newClientConfig())
	if err != nil {
		return nil, fmt.Errorf("failed to create pubsub client: %w", err)
	}

	return &Client{
		ProjectId: projectId,
		Client:    client,
	}, nil
}

// Close gracefully shuts down the Pub/Sub client, freeing any resources.
// Returns any error that occurs while closing the client.
func (c *Client) Close() error {
	return c.Client.Close()
}
