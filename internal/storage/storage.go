package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudevents/sdk-go/v2/event"
)

type StorageObjectData struct {
	Bucket         string    `json:"bucket,omitempty"`
	Name           string    `json:"name,omitempty"`
	Metageneration int64     `json:"metageneration,string,omitempty"`
	TimeCreated    time.Time `json:"timeCreated,omitempty"`
	Updated        time.Time `json:"updated,omitempty"`
}

func ProcessFile(ctx context.Context, event event.Event) error {
	var data StorageObjectData

	if err := event.DataAs(&data); err != nil {
		return fmt.Errorf("failed to parse event data: %v", err)
	}

	log.Printf("Attempting to process file: %v", data.Name)

	return nil
}
