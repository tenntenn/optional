package pattern_test

import (
	"regexp"
	"testing"

	"github.com/tenntenn/optional"
	. "github.com/tenntenn/optional/pattern"
)

var (
	reg = regexp.MustCompile("^go*$")
)

func TestPattern(t *testing.T) {
	data := []struct {
		P         *Pattern
		V         string
		IsMatched bool
	}{
		{P: WithString("test"), V: "test", IsMatched: true},
		{P: WithString("test"), V: "foo", IsMatched: false},
		{P: WithRegexp(reg), V: "go", IsMatched: true},
		{P: WithRegexp(reg), V: "goooo", IsMatched: true},
		{P: WithRegexp(reg), V: "goto", IsMatched: false},
		{P: &Pattern{S: optional.NewString("test"), R: reg}, V: "test", IsMatched: true},
		{P: &Pattern{S: optional.NewString("test"), R: reg}, V: "go", IsMatched: false},
		{P: &Pattern{}, V: "test", IsMatched: false},
		{P: nil, V: "test", IsMatched: false},
	}

	for i, d := range data {
		if isMatched := d.P.Match(d.V); isMatched != d.IsMatched {
			t.Errorf("data[%d].P.Matched expects %v, but %v", i, d.IsMatched, isMatched)
		}
	}
}
