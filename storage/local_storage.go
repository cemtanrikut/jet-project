package storage

import (
	"os"
)

// LocalStorageStrategy implements the StorageStrategy interface for local storage.
type LocalStorageStrategy struct{}

func (l *LocalStorageStrategy) SaveContent(content string, fileName string) error {
	return os.WriteFile(fileName, []byte(content), 0644)
}
