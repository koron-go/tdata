# tdata - testdata load utility

[![GoDoc](https://godoc.org/github.com/koron-go/tdata?status.svg)](https://godoc.org/github.com/koron-go/tdata)
[![CircleCI](https://circleci.com/gh/koron-go/tdata.svg?style=svg)](https://circleci.com/gh/koron-go/tdata)

tdata and related pacakges help to access files in testdata/ directory for
test.

## Examples

```golang
import "github.com/koron-go/tdata"

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

content of testdata/withjson.json:

```json
[
  { "foo": "abc", "bar": 123 },
  { "foo": "xyz", "bar": 999 }
]
```
