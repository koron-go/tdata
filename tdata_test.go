package tdata

import (
	"reflect"
	"testing"
)

func TestNorm(t *testing.T) {
	for _, d := range []struct{ in, exp string }{
		{"TestWithJSON", "withjson"},
		{"BenchWithJSON", "benchwithjson"},
		{"a/b/c", "a/b/c"},
		{"TestFoo/Bar/Baz", "foo/bar/baz"},
		{"TestFoo/Bar Baz", "foo/bar_baz"},
		{"TestFoo/Bar  Baz", "foo/bar_baz"},
	} {
		act := norm(d.in)
		if act != d.exp {
			t.Errorf("failed: in=%q exp=%q act=%q", d.in, d.exp, act)
		}
	}
}

type jsonData struct {
	FooJ string `json:"foo"`
	BarJ int    `json:"bar"`
	BazJ bool   `json:"baz"`
}

func TestSimpleJSON(t *testing.T) {
	var v []jsonData
	New(t).JSON().Decode(&v)
	if !reflect.DeepEqual(v, []jsonData{
		{"jsonABC", 123, true},
		{"jsonXYZ", 999, true},
		{"jsonABC", 123, false},
	}) {
		t.Errorf("not match: %+v", v)
	}
}

type yamlData struct {
	FooY string `yaml:"foo"`
	BarY int    `yaml:"bar"`
	BazY bool   `yaml:"baz"`
}

func TestSimpleYAML(t *testing.T) {
	var v []yamlData
	New(t).YAML().Decode(&v)
	if !reflect.DeepEqual(v, []yamlData{
		{"yamlABC", 123, true},
		{"yamlXYZ", 999, true},
		{"yamlABC", 123, false},
	}) {
		t.Errorf("not match: %+v", v)
	}
}

func TestThroughJSON(t *testing.T) {
	var v []jsonData
	New(t).YAML().JSON().Decode(&v)
	if !reflect.DeepEqual(v, []jsonData{
		{"yaml+json_ABC", 123, true},
		{"yaml+json_XYZ", 999, true},
		{"yaml+json_ABC", 123, false},
	}) {
		t.Errorf("not match: %+v", v)
	}
}

var g = struct {
	Value1 int
	Value2 int
	Value3 int
}{
	Value1: 987,
	Value2: 654,
	Value3: 321,
}

func TestTemplateYAML(t *testing.T) {
	var v []yamlData
	New(t).Template(g).YAML().Decode(&v)
	if !reflect.DeepEqual(v, []yamlData{
		{"yamlABC", 987, true},
		{"yamlXYZ", 654, true},
		{"yamlABC", 321, false},
	}) {
		t.Errorf("not match: %+v", v)
	}
}
