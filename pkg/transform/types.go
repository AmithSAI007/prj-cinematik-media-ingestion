package transform

import (
	"time"
)

// StorageObjectData represents the metadata of a storage object (e.g., a file)
// stored in a cloud storage bucket. It includes information like the bucket
// name, object name, metadata generation, creation time, and last update time.
type StorageObjectData struct {
	Bucket         string    `json:"bucket,omitempty"`
	Name           string    `json:"name,omitempty"`
	Metageneration string    `json:"metageneration,omitempty"`
	TimeCreated    time.Time `json:"timecreated,omitempty"`
	Updated        time.Time `json:"updated,omitempty"`
	ContentType    string    `json:"contentType"`
	Size           string    `json:"size"`
}

// TopicMessgaeData represents the metadata of a file related to a topic message.
// It includes file details like the bucket, file name, file path, content type,
// size, and the time it was created.
type TopicMessgaeData struct {
	Bucket      string    `json:"bucket,omitempty"`
	FileName    string    `json:"filename,omitempty"`
	FilePath    string    `json:"filepath,omitempty"`
	ContentType string    `json:"contenttype,omitempty"`
	Size        string    `json:"size,omitempty"`
	TimeCreated time.Time `json:"timecreated,omitempty"`
}
