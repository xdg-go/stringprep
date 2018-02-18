package saslprep

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/xdg/stringprep"
)

type profileCase struct {
	label string
	in    string
	out   string
	err   error
}

var profileTests = []profileCase{
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
		err:   stringprep.ErrProhibited(fmt.Errorf("character '0x%04x' is prohibited", 0x0007)),
	},
	{
		label: "bidi not ok",
		in:    "\u0627\u0031",
		out:   "",
		err:   stringprep.ErrBiDi(fmt.Errorf("Invalid use of bidirectional characters")),
	},
}

func TestSASLprep(t *testing.T) {
	for _, c := range profileTests {
		t.Run(c.label, func(t *testing.T) {
			got, err := Prepare(c.in)
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
