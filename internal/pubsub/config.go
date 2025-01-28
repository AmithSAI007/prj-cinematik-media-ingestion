package pubsub

import (
	"time"

	"cloud.google.com/go/pubsub"
	vkit "cloud.google.com/go/pubsub/apiv1"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/grpc/codes"
)

// newClientConfig returns a new Pub/Sub client configuration with custom retry settings
// for publishing messages. The configuration sets up a retry policy for failed publish
// operations with specified retryable error codes and backoff settings.
func newClientConfig() *pubsub.ClientConfig {
	return &pubsub.ClientConfig{
		PublisherCallOptions: &vkit.PublisherCallOptions{
			Publish: []gax.CallOption{
				gax.WithRetry(func() gax.Retryer {
					return gax.OnCodes([]codes.Code{
						codes.Aborted,
						codes.Canceled,
						codes.Internal,
						codes.ResourceExhausted,
						codes.Unknown,
						codes.Unavailable,
						codes.DeadlineExceeded,
					}, gax.Backoff{
						Initial:    250 * time.Millisecond,
						Max:        60 * time.Second,
						Multiplier: 1.45,
					})
				}),
			},
		},
	}
}
