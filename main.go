package main

import (
	"fmt"
	"log"

	"github.com/cemtanrikut/jet-project/token"

	"github.com/cemtanrikut/jet-project/storage"
)

func main() {
	// Initialize Token Provider
	tokenProvider := &token.SalesforceTokenProvider{
		ClientID:     "client-id",
		ClientSecret: "client-secret",
		AuthURL:      "https://auth.salesforce.com/token",
	}

	accessToken, err := tokenProvider.GetAccessToken()
	if err != nil {
		log.Fatalf("Failed to get access token: %v", err)
	}

	fmt.Printf("Access Token: %s\n", accessToken)

	// Select Storage Strategy
	var storageStrategy storage.StorageStrategy
	storageType := "local" // Can be "local", "s3", "gcs"

	switch storageType {
	case "local":
		storageStrategy = &storage.LocalStorageStrategy{}
	case "s3":
		storageStrategy = &storage.S3StorageStrategy{}
	case "gcs":
		storageStrategy = &storage.GCSStorageStrategy{}
	default:
		log.Fatalf("Invalid storage type: %s", storageType)
	}

	// Save Content
	content := `{"example": "This is a test content block"}`
	err = storageStrategy.SaveContent(content, "output.json")
	if err != nil {
		log.Fatalf("Failed to save content: %v", err)
	}

	fmt.Println("Content saved successfully!")
}
