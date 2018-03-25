package tdata

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/koron-go/tdata"
	yaml "gopkg.in/yaml.v2"
)

// Unmarshal decodes a testdata file (.yml) as YAML.
func Unmarshal(t *testing.T, v interface{}) {
	t.Helper()
	f := tdata.Open(t, ".yml")
	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("failed to load testdata for YAML: %s", err)
		f.Close()
	}
	f.Close()
	err = yaml.Unmarshal(b, &v)
	if err != nil {
		t.Fatalf("failed to decode testdata as YAML: %s", err)
	}
}

// UnmarshalJSON loads a testdata file (.yml) as YAML, converts it to JSON then
// decodes it.
func UnmarshalJSON(t *testing.T, v interface{}) {
	t.Helper()
	var v0 interface{}
	Unmarshal(t, &v0)
	b, err := json.Marshal(v0)
	if err != nil {
		t.Fatalf("failed to convert to JSON: %s", err)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Fatalf("failed to decode as JSON: %s", err)
	}
}