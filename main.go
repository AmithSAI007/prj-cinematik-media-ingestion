package main

import (
	"context"
	"log"

	"github.com/AmithSAI007/prj-cinematik-media-ingestion/internal/storage"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

func init() {
	functions.CloudEvent("MediaIngestionService", mediaIngestionService)
}

type Metadata struct {
	VideoId    string  `json:"video_id"`
	FileName   string  `json:"file_name"`
	BucketPath string  `json:"bucket_path"`
	Duration   float64 `json:"duration"`
	Bitrate    int64   `json:"bitrate"`
	Resolution string  `json:"resolution"`
	FrameRate  string  `json:"frame_rate"`
	Codec      string  `json:"codec"`
	UploadTime string  `json:"upload_time"`
}

func mediaIngestionService(ctx context.Context, event event.Event) error {
	if err := storage.ProcessFile(ctx, event); err != nil {
		log.Printf("Error processing file: %v", err)
		return err
	}
	return nil
}
