package transform

import (
	"fmt"
)

// TransformToTopicMessageData converts a StorageObjectData object to a TopicMessgaeData object.
// It uses the information from the StorageObjectData, such as the bucket, name, and time created,
// to populate the corresponding fields in the TopicMessgaeData. The ContentType and Size fields are
// left empty or set to zero as they are not available in the input data.
func TransformToTopicMessageData(data StorageObjectData) (TopicMessgaeData, error) {
	if data.Bucket == "" || data.Name == "" {
		return TopicMessgaeData{}, fmt.Errorf("invalid input: bucket or name is empty")
	}
	return TopicMessgaeData{
		Bucket:      data.Bucket,
		FileName:    data.Name,
		FilePath:    fmt.Sprintf("gs://%s/%s", data.Bucket, data.Name),
		ContentType: data.ContentType,
		Size:        data.Size,
		TimeCreated: data.TimeCreated,
	}, nil
}
