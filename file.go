package tdata

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Open opens a testdata file for the test.
func Open(t *testing.T, ext string) *os.File {
	t.Helper()
	f, err := os.Open(name(t, ext))
	if err != nil {
		t.Fatalf("failed to open testdata: %s", err)
	}
	return f
}

var nrep = strings.NewReplacer("/", "-", "\\", "-", " ", "_")

func normalize(s string) string {
	return nrep.Replace(strings.ToLower(strings.TrimPrefix(s, "Test")))
}

func name(t *testing.T, ext string) string {
	return filepath.Join("testdata", normalize(t.Name())) + ext
}
