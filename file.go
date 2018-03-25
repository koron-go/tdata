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
	n := filepath.Join("testdata", normalize(t.Name()))
	f, err := os.Open(n + ext)
	if err != nil {
		t.Fatalf("failed to open testdata: %s", err)
	}
	return f
}

var nrep = strings.NewReplacer("/", "-", "\\", "-", " ", "_")

func normalize(s string) string {
	return nrep.Replace(strings.ToLower(strings.TrimPrefix(s, "Test")))
}
