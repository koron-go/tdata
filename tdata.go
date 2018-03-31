package tdata

import (
	"bytes"
	encjson "encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"text/template"

	encyaml "gopkg.in/yaml.v2"
)

// New creates a testdata loader.
func New(tb testing.TB) Loader {
	return &tdata{
		tb:  tb,
		nam: "testdata",
	}
}

// NewName creates a file loader.
func NewName(name string) Loader {
	return &tdata{
		nam: name,
	}
}

// Loader is a testdata/file loader.
type Loader interface {
	Template(d interface{}) TemplateLoader
	YAML() YAMLDecoder
	JSON() JSONDecoder
}

// TemplateLoader is a testdata/file loader with template.
type TemplateLoader interface {
	YAML() YAMLDecoder
	JSON() JSONDecoder
}

// YAMLDecoder decodes a testdata/file as YAML.
type YAMLDecoder interface {
	Decode(v interface{})
	JSON() JSONDecoder
}

// JSONDecoder decodes a testdata/file as JSON.
type JSONDecoder interface {
	Decode(v interface{})
}

type tdata struct {
	tb  testing.TB
	nam string

	mode    tmode
	revExts []string
	tmpl    bool
	data    interface{}
	yaml    bool
}

func (td *tdata) Template(data interface{}) TemplateLoader {
	if td.mode != none {
		panic(fmt.Sprintf("unexpected mode in Template: %v", td.mode))
	}
	td.mode = tmpl
	td.revExts = append(td.revExts, ".tmpl")
	td.tmpl = true
	td.data = data
	return td
}

func (td *tdata) YAML() YAMLDecoder {
	if td.mode != none && td.mode != tmpl {
		panic(fmt.Sprintf("unexpected mode in YAML: %v", td.mode))
	}
	td.mode = yaml
	td.revExts = append(td.revExts, ".yaml")
	td.yaml = true
	return td
}

func (td *tdata) JSON() JSONDecoder {
	if td.mode != none && td.mode != tmpl && td.mode != yaml {
		panic(fmt.Sprintf("unexpected mode in JSON: %v", td.mode))
	}
	td.mode = json
	td.revExts = append(td.revExts, ".json")
	return td
}

func (td *tdata) Decode(v interface{}) {
	if td.mode != yaml && td.mode != json {
		panic(fmt.Sprintf("unexpected mode in Decode: %v", td.mode))
	}
	mode := td.mode
	td.mode = fin
	if td.tb != nil {
		td.tb.Helper()
	}

	b, err := td.load()
	if err != nil {
		td.fail("failed to open testdata: %s", err)
	}

	if mode == yaml {
		err := encyaml.Unmarshal(b, v)
		if err != nil {
			td.fail("failed to decode testdata as YAML: %s", err)
		}
		return
	}

	if td.yaml {
		var v0 interface{}
		err := encyaml.Unmarshal(b, &v0)
		if err != nil {
			td.fail("failed to decode testdata as intermediate YAML: %s", err)
		}
		b, err = encjson.Marshal(v0)
		if err != nil {
			td.fail("failed to convert to JSON: %s", err)
		}
	}

	err = encjson.Unmarshal(b, v)
	if err != nil {
		td.fail("failed to decode as JSON: %s", err)
	}
}

func (td *tdata) fail(s string, args ...interface{}) {
	if td.tb == nil {
		panic(fmt.Sprintf(s, args...))
	}
	td.tb.Helper()
	td.tb.Fatalf(s, args...)
}

func (td *tdata) name() string {
	sb := strings.Builder{}
	if td.tb != nil {
		sb.WriteString(filepath.Join(td.nam, norm(td.tb.Name())))
	} else {
		sb.WriteString(td.nam)
	}
	for i := len(td.revExts) - 1; i >= 0; i-- {
		sb.WriteString(td.revExts[i])
	}
	return sb.String()
}

func (td *tdata) load() ([]byte, error) {
	fname := td.name()
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	if !td.tmpl {
		return b, nil
	}

	tmpl, err := template.New(fname).Parse(string(b))
	if err != nil {
		return nil, err
	}
	bb := &bytes.Buffer{}
	err = tmpl.Execute(bb, td.data)
	if err != nil {
		return nil, err
	}
	return bb.Bytes(), nil
}

type tmode int

const (
	none tmode = iota
	tmpl
	yaml
	json
	fin
)

var notAlnumSlash = regexp.MustCompile(`[^0-9A-Za-z/]+`)

func norm(s string) string {
	s = strings.TrimPrefix(s, "Test")
	s = notAlnumSlash.ReplaceAllString(s, "_")
	s = strings.ToLower(s)
	return s
}
