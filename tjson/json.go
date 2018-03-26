package tjson

import (
	"encoding/json"
	"testing"

	"github.com/koron-go/tdata"
)

// Unmarshal decodes a testdata file (.json) as JSON.
func Unmarshal(t *testing.T, v interface{}) {
	t.Helper()
	f := tdata.Open(t, ".json")
	defer f.Close()
	d := json.NewDecoder(f)
	err := d.Decode(&v)
	if err != nil {
		t.Fatalf("failed to decode testdata as JSON: %s", err)
	}
}
