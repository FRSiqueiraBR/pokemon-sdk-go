package pokemon_test

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestApplicationLayerDoesNotImportInfra(t *testing.T) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("could not get caller path")
	}

	root := filepath.Clean(filepath.Join(filepath.Dir(file), "..", "..", ".."))
	appDir := filepath.Join(root, "internal", "application")

	err := filepath.WalkDir(appDir, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if strings.Contains(string(content), "internal/infra") {
			t.Fatalf("forbidden import found in %s", path)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("scan failed: %v", err)
	}
}

func TestPublicAPIHidesHTTPClientEntryPoint(t *testing.T) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("could not get caller path")
	}

	root := filepath.Clean(filepath.Join(filepath.Dir(file), "..", "..", ".."))
	pkgDir := filepath.Join(root, "pkg", "pokemon")

	err := filepath.WalkDir(pkgDir, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		src := string(content)
		if strings.Contains(src, "type Client struct") || strings.Contains(src, "func NewClient(") {
			t.Fatalf("forbidden public client symbol found in %s", path)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("scan failed: %v", err)
	}
}
