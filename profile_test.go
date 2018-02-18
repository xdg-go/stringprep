package stringprep

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

var p1 = Profile{
	Mappings:  []Mapping{TableB1, TableB2},
	Normalize: true,
	Prohibits: []Set{TableC1_1, TableC1_2, TableC6},
	CheckBiDi: true,
}

var p2 = Profile{
	Mappings:  []Mapping{},
	Normalize: false,
	Prohibits: []Set{},
	CheckBiDi: false,
}

type profileCase struct {
	label   string
	profile Profile
	in      string
	out     string
	err     error
}

var profileTests = []profileCase{
	{label: "p1: empty", profile: p1, in: "", out: "", err: nil},
	{label: "p1: ascii", profile: p1, in: "user", out: "user", err: nil},
	{label: "p1: zws", profile: p1, in: "u\u200Ber", out: "uer", err: nil},
	{label: "p1: sharp", profile: p1, in: "u\u00DFer", out: "usser", err: nil},
	{label: "p1: nfkc", profile: p1, in: "ua\u030Aer", out: "u\u00e5er", err: nil},
	{
		label:   "p1: replacement",
		profile: p1,
		in:      "u\uFFFDer",
		out:     "",
		err:     ErrProhibited(fmt.Errorf("character '0x%04x' is prohibited", 0xFFFD)),
	},
	{
		label:   "p1: bidi ok",
		profile: p1,
		in:      "\u0627\u0031\u0628",
		out:     "\u0627\u0031\u0628",
		err:     nil,
	},
	{
		label:   "p1: bidi not ok; RandAL + L",
		profile: p1,
		in:      "\u0627\u0589\u0628",
		out:     "",
		err:     ErrBiDi(fmt.Errorf("Invalid use of bidirectional characters")),
	},
	{
		label:   "p1: bidi not ok; endpoints",
		profile: p1,
		in:      "\u0627\u0031",
		out:     "",
		err:     ErrBiDi(fmt.Errorf("Invalid use of bidirectional characters")),
	},
	{
		label:   "p1: bidi not ok; chars",
		profile: p1,
		in:      "\u206D",
		out:     "",
		err:     ErrBiDi(fmt.Errorf("Invalid use of bidirectional characters")),
	},
	{
		label:   "p2: bidi not checked",
		profile: p2,
		in:      "\u0627\u0031",
		out:     "\u0627\u0031",
		err:     nil,
	},
	{label: "p2: no nfkc", profile: p2, in: "ua\u030Aer", out: "ua\u030Aer", err: nil},
}

func TestProfiles(t *testing.T) {
	for _, c := range profileTests {
		t.Run(c.label, func(t *testing.T) {
			got, err := c.profile.Prepare(c.in)
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
