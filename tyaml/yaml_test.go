package tyaml

import (
	"reflect"
	"testing"
)

type Data struct {
	Foo string `yaml:"foo"`
	Bar int    `yaml:"bar"`
}

func TestUnmarshalSingle(t *testing.T) {
	var d Data
	// testdata/unmarshalsingle.yml
	Unmarshal(t, &d)
	if !reflect.DeepEqual(d, Data{"abc", 123}) {
		t.Errorf("unexpected JSONData: %+v", d)
	}
}

func TestUnmarshalArray(t *testing.T) {
	var d []Data
	// testdata/unmarshalarray.yml
	Unmarshal(t, &d)
	if !reflect.DeepEqual(d, []Data{
		{"abc", 123},
		{"xyz", 999},
		{"foo", 600},
	}) {
		t.Errorf("unexpected array: %+v", d)
	}
}

type JSONData struct {
	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}

func TestUnmarshalJSONSingle(t *testing.T) {
	var d JSONData
	// testdata/unmarshaljsonsingle.yml
	Unmarshal(t, &d)
	if !reflect.DeepEqual(d, JSONData{"abc", 123}) {
		t.Errorf("unexpected JSONData: %+v", d)
	}
}

func TestUnmarshalJSONArray(t *testing.T) {
	var d []JSONData
	// testdata/unmarshaljsonarray.yml
	Unmarshal(t, &d)
	if !reflect.DeepEqual(d, []JSONData{
		{"abc", 123},
		{"xyz", 999},
		{"foo", 600},
	}) {
		t.Errorf("unexpected array: %+v", d)
	}
}
