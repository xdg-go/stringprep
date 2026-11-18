package stringprep

import (
	"strings"
	"testing"
)

func FuzzSASLprepIdempotent(f *testing.F) {
	seeds := []string{
		"",
		"user",
		"I\u00ADX",
		"\u00AA",
		"\u2168",
		"a\u00A0b",
		"\u0627\u0031\u0628",
		// GH issue #7
		strings.Repeat("\u0344", 9) + strings.Repeat("\u0345", 13),
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, in string) {
		once, err := SASLprep.Prepare(in)
		if err != nil {
			return
		}
		twice, err := SASLprep.Prepare(once)
		if err != nil {
			t.Fatalf("Prepare(%+q) = %+q, but re-preparing it fails: %v", in, once, err)
		}
		if twice != once {
			t.Fatalf("not idempotent for %+q:\nonce:  %+q\ntwice: %+q", in, once, twice)
		}
	})
}
