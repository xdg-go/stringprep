package stringprep

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestSASLprep(t *testing.T) {
	saslTests := []struct {
		label string
		in    string
		out   string
		err   error
	}{
		{label: "soft hyphen", in: "I\u00ADX", out: "IX", err: nil},
		{label: "non ASCII space", in: "I\u2000X", out: "I X", err: nil},
		{label: "no transform", in: "user", out: "user", err: nil},
		{label: "case preserve", in: "USER", out: "USER", err: nil},
		{label: "8859-1 to NFKC", in: "\u00AA", out: "a", err: nil},
		{label: "NFKC", in: "\u2168", out: "IX", err: nil},
		{
			label: "prohibited",
			in:    "\u0007",
			out:   "",
			err:   Error{Msg: errProhibited, Rune: '\u0007'},
		},
		{
			label: "bidi not ok",
			in:    "\u0627\u0031",
			out:   "",
			err:   Error{Msg: errLastRune, Rune: '\u0031'},
		},
		{
			label: "30 combining marks",
			in:    "a" + strings.Repeat("\u0301", 30),
			out:   "\u00E1" + strings.Repeat("\u0301", 29),
			err:   nil,
		},
		{
			label: "31 combining marks",
			in:    "a" + strings.Repeat("\u0301", 31),
			out:   "",
			err:   Error{Msg: errTooManyCombining, Rune: '\u034F'},
		},
		{
			label: "precomposed counts its marks",
			in:    "\u00E1" + strings.Repeat("\u0301", 30),
			out:   "",
			err:   Error{Msg: errTooManyCombining, Rune: '\u034F'},
		},
		{
			label: "base character resets count",
			in:    "a" + strings.Repeat("\u0301", 30) + "b" + strings.Repeat("\u0301", 30),
			out:   "\u00E1" + strings.Repeat("\u0301", 29) + "b" + strings.Repeat("\u0301", 30),
			err:   nil,
		},
		{
			label: "GH issue 7",
			in:    strings.Repeat("\u0344", 9) + strings.Repeat("\u0345", 13),
			out:   "",
			err:   Error{Msg: errTooManyCombining, Rune: '\u034F'},
		},
	}

	for _, c := range saslTests {
		t.Run(c.label, func(t *testing.T) {
			got, err := SASLprep.Prepare(c.in)
			t.Logf("err is '%v'", err)
			if c.err == nil {
				if got != c.out {
					t.Errorf("input '%s': got '%s', expected '%s'",
						strconv.QuoteToASCII(c.in),
						strconv.QuoteToASCII(got),
						strconv.QuoteToASCII(c.out))
				}
			} else {
				if !reflect.DeepEqual(err, c.err) {
					t.Errorf("input '%s': got error '%v', expected '%v'",
						strconv.QuoteToASCII(c.in), err, c.err)
				}
			}

		})
	}
}
