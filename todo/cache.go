package todo

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Cached variables.
type Vars struct {
	SelectedTodoID string `yaml:"selectedTodoID"`
}

var CacheFilePath = getCacheFilePath()

func getCacheFilePath() string {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		// TODO: Do some proper error handling and have a sane fallback
		cacheDir = "."
	}

	return filepath.Join(cacheDir, "todogo", "cache.yml")
}

func InitCache() error {
	fileInfo, err := os.Stat(CacheFilePath)

	if fileInfo != nil {
		return nil
	}

	// We simply ignore the ErrNotExist and handle other errors
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("error checking %s: %w", CacheFilePath, err)
	}

	if err := os.MkdirAll(filepath.Dir(CacheFilePath), 0770); err != nil {
		return fmt.Errorf("error creating cache directory for %s: %w", CacheFilePath, err)
	}

	defaultVars := Vars{
		SelectedTodoID: "",
	}

	varsB, err := yaml.Marshal(defaultVars)
	if err != nil {
		return fmt.Errorf("failed to marshal default cache vars: %w", err)
	}

	f, err := os.Create(CacheFilePath)
	if err != nil {
		return fmt.Errorf("failed to create %s: %w", CacheFilePath, err)
	}
	defer f.Close()

	if _, err := f.Write(varsB); err != nil {
		return fmt.Errorf("failed to write to %s: %w", CacheFilePath, err)
	}

	return nil
}
