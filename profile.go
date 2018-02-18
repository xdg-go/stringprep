package stringprep

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

// ErrBiDi indicates invalid use of bidirectional characters.
type ErrBiDi error

// ErrProhibited indicates a prohibited character was found.
type ErrProhibited error

// Profile represents a stringprep profile.
type Profile struct {
	Mappings  []Mapping
	Normalize bool
	Prohibits []Set
	CheckBiDi bool
}

// Prepare transforms an input string to an output string following
// the rules defined in the profile as defined by RFC-3454.
func (p Profile) Prepare(s string) (string, error) {
	// Optimistically, assume output will be same length as input
	temp := make([]rune, 0, len(s))

	// Apply maps
	for _, r := range s {
		rs, ok := p.applyMaps(r)
		if ok {
			temp = append(temp, rs...)
		} else {
			temp = append(temp, r)
		}
	}

	// Normalize
	var out string
	if p.Normalize {
		out = norm.NFKC.String(string(temp))
	} else {
		out = string(temp)
	}

	// Check prohibited
	for _, r := range out {
		if p.runeIsProhibited(r) {
			return "", ErrProhibited(fmt.Errorf("character '0x%04x' is prohibited", r))
		}
	}

	// Check BiDi allowed
	if p.CheckBiDi {
		if !passesBiDiRules(out) {
			return "", ErrBiDi(fmt.Errorf("Invalid use of bidirectional characters"))
		}
	}

	return out, nil
}

func (p Profile) applyMaps(r rune) ([]rune, bool) {
	for _, m := range p.Mappings {
		rs, ok := m.Map(r)
		if ok {
			return rs, true
		}
	}
	return nil, false
}

func (p Profile) runeIsProhibited(r rune) bool {
	for _, s := range p.Prohibits {
		if s.Contains(r) {
			return true
		}
	}
	return false
}
