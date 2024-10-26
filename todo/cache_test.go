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

func TestSetVar(t *testing.T) {
	CacheFilePath := setupTempCache(t)

	if err := InitCache(); err != nil {
		t.Fatalf("Failed to initialize cache: %v", err)
	}

	wantID := "test-id"
	if err := SetVar("SelectedTodoID", wantID); err != nil {
		t.Fatalf("SetVar returned an error: %v", err)
	}

	var vars Vars

	varsB, err := os.ReadFile(CacheFilePath)
	if err != nil {
		t.Fatalf("Failed to read cache file: %v", err)
	}

	if err := yaml.Unmarshal(varsB, &vars); err != nil {
		t.Fatalf("Failed to unmarshal cache vars: %v", err)
	}

	if vars.SelectedTodoID != wantID {
		t.Errorf("want SelectedTodoID '%s', got '%s'", wantID, vars.SelectedTodoID)
	}

	err = SetVar("UnknownKey", "some value")
	if err == nil || err.Error() != "unknown variable name: UnknownKey" {
		t.Errorf("want error for unknown variable, got: %v", err)
	}
}
