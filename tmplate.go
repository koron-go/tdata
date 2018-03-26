package tdata

import (
	"bytes"
	"io"
	"testing"
	"text/template"
)

// TemplateData is used for TemplateExecute.
var TemplateData interface{}

// ExecuteText loads a testdata file as template and expand/execute it.
func ExecuteText(t *testing.T, ext string) io.Reader {
	t.Helper()
	tmpl, err := template.ParseFiles(name(t, ext))
	if err != nil {
		t.Fatalf("failed to parse template: %s", err)
	}
	bb := &bytes.Buffer{}
	err = tmpl.Execute(bb, TemplateData)
	if err != nil {
		t.Fatalf("failed to execute template: %s", err)
	}
	return bb
}
