package storage

import (
	"os"
	"testing"
)

func TestLocalStorageStrategy_SaveContent(t *testing.T) {
	strategy := &LocalStorageStrategy{}
	fileName := "test_output.json"
	content := `{"key": "value"}`

	// Test saving content to local storage
	err := strategy.SaveContent(content, fileName)
	if err != nil {
		t.Errorf("Failed to save content to local storage: %v", err)
	}

	// Verify the file was created
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Errorf("File not created: %s", fileName)
	}

	// Cleanup
	_ = os.Remove(fileName)
}

func TestS3StorageStrategy_SaveContent(t *testing.T) {
	strategy := &S3StorageStrategy{}
	err := strategy.SaveContent("test content", "test_file.json")
	if err == nil {
		t.Errorf("Expected error for unimplemented S3 logic, but got nil")
	}
}

func TestGCSStorageStrategy_SaveContent(t *testing.T) {
	strategy := &GCSStorageStrategy{}
	err := strategy.SaveContent("test content", "test_file.json")
	if err == nil {
		t.Errorf("Expected error for unimplemented GCS logic, but got nil")
	}
}
