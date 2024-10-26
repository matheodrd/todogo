package todo

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
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

func TestLoadCache(t *testing.T) {
	CacheFilePath := setupTempCache(t)
	wantVars := Vars{
		SelectedTodoID: "test-id",
	}

	varsB, err := yaml.Marshal(&wantVars)
	if err != nil {
		t.Fatalf("Failed to marshal cache vars: %v", err)
	}

	if err := os.WriteFile(CacheFilePath, varsB, 0600); err != nil {
		t.Fatalf("Failed to write vars to cache file: %v", err)
	}

	gotVars, err := LoadCache()
	if err != nil {
		t.Fatalf("LoadCache returned an error: %v", err)
	}

	if wantVars.SelectedTodoID != gotVars.SelectedTodoID {
		t.Errorf("want SelectedTodoID '%s', got '%s'", wantVars.SelectedTodoID, gotVars.SelectedTodoID)
	}
}
