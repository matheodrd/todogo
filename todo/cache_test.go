package todo

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func setupTempCache(t *testing.T) string {
	t.Helper()
	tmpDir := t.TempDir()
	CacheFilePath = filepath.Join(tmpDir, "cache.yml")
	return CacheFilePath
}

func TestInitCache(t *testing.T) {
	cacheFilePath := setupTempCache(t)
	wantCacheVars := "selectedTodoID: \"\""

	if err := InitCache(); err != nil {
		t.Fatalf("Failed to init cache file: %v", err)
	}

	gotCacheVarsB, err := os.ReadFile(cacheFilePath)
	if err != nil {
		t.Fatalf("Failed to read cache file: %v", err)
	}
	if !strings.Contains(string(gotCacheVarsB), wantCacheVars) {
		t.Errorf("Cache file does not contain want vars: want '%s', got '%s'", wantCacheVars, gotCacheVarsB)
	}
}
