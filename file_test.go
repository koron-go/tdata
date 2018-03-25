package tdata

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	for _, d := range []struct{ in, exp string }{
		{"a/b/c", "a-b-c"},
		{"TestFoo/Bar/Baz", "foo-bar-baz"},
	} {
		act := normalize(d.in)
		if act != d.exp {
			t.Errorf("failed: in=%q exp=%q act=%q", d.in, d.exp, act)
		}
	}
}
