# tdata - testdata load utility

tdata and related pacakges help to access files in testdata/ directory for
test.

## Examples

```golang
import "github.com/koron-go/tdata/tjson"

type data struct {
	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}

func TestWithJSON(t *testing.T) {
	var v []data
	tjson.Unmarshal(t, &v)
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
