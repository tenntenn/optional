package pattern

import (
	"regexp"

	"github.com/tenntenn/optional"
)

// Parttern is represented by string value or Regexp.
// It can test whether matched with string values or not.
type Parttern struct {
	S *optional.String
	R *regexp.Regexp
}

func WithString(s string) *Parttern {
	return &Parttern{S: optional.NewString(s)}
}

func WithRegexp(r *regexp.Regexp) *Parttern {
	return &Parttern{R: r}
}

// Match tests whether s is matched with the pattern.
// if p.S is set, p.S is used to test for matching.
// if p.S is nil and p.R is set, p.R is used to test.
// if p is nil Match returns false.
func (p *Parttern) Match(s string) bool {
	if p == nil {
		return false
	}

	switch {
	case p.S != nil:
		return p.S.IsSet(s)
	case p.R != nil:
		return p.R.MatchString(s)
	}
	return false
}
