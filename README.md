# tdata - testdata load utility

[![GoDoc](https://godoc.org/github.com/koron-go/tdata?status.svg)](https://godoc.org/github.com/koron-go/tdata)
[![CircleCI](https://circleci.com/gh/koron-go/tdata.svg?style=svg)](https://circleci.com/gh/koron-go/tdata)

tdata helps to access files in testdata/ directory for test.

## Getting Started

This example load a JSON file `testdata/withjson.json`.
The name of the file is determined from `t.Name()` automatically.

```golang
import (
	"testing"
	"github.com/koron-go/tdata"
)

type data struct {
	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}

func TestWithJSON(t *testing.T) {
	var v []data
	tdata.New(t).JSON().Decode(&v)
	// use v as test data: load from "testdata/withjson.json"
}
```

Contents example of testdata/withjson.json:

```json
[
  { "foo": "abc", "bar": 123 },
  { "foo": "xyz", "bar": 999 }
]
```

### To load a YAML

When you want to decode YAML, replace `Decode()` line like this:

```golang
	tdata.New(t).YAML().Decode(&v)
```

This loads a file `testdata/withjson.yaml`.
Thus, extension of filename is determined by intermediate of chain.

### Convert YAML to JSON

This is useful when destination structure doesn't have any `yaml` tags, but you
want to write data in YAML format.

```golang
	tdata.New(t).YAML().JSON().Decode(&v)
```

This loads a file `testdata/withjson.json.yaml`.

### Templating

```golang
	tdata.New(t).Template(g).YAML().JSON().Decode(&v)
```

This loads a file `testdata/withjson.json.yaml.tmpl`.

See [text/template](https://golang.org/pkg/text/template/) for syntax of
template.
