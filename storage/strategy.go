package storage

// StorageStrategy defines the interface for different storage strategies.
type StorageStrategy interface {
	SaveContent(content string, fileName string) error
}
