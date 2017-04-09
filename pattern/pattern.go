package pattern

import (
	"regexp"

	"github.com/tenntenn/optional"
)

// Pattern is represented by string value or Regexp.
// It can test whether matched with string values or not.
type Pattern struct {
	S *optional.String
	R *regexp.Regexp
}

func WithString(s string) *Pattern {
	return &Pattern{S: optional.NewString(s)}
}

func WithRegexp(r *regexp.Regexp) *Pattern {
	return &Pattern{R: r}
}

// Match tests whether s is matched with the pattern.
// if p.S is set, p.S is used to test for matching.
// if p.S is nil and p.R is set, p.R is used to test.
// if p is nil Match returns false.
func (p *Pattern) Match(s string) bool {
	if p == nil {
		return false
	}

	switch {
	case p.S != nil:
		return p.S.Match(s)
	case p.R != nil:
		return p.R.MatchString(s)
	}
	return false
}
