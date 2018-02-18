package saslprep

import "github.com/xdg/stringprep"

var mapNonASCIISpaceToASCIISpace = stringprep.Mapping{
	0x00A0: []rune{0x0020},
	0x1680: []rune{0x0020},
	0x2000: []rune{0x0020},
	0x2001: []rune{0x0020},
	0x2002: []rune{0x0020},
	0x2003: []rune{0x0020},
	0x2004: []rune{0x0020},
	0x2005: []rune{0x0020},
	0x2006: []rune{0x0020},
	0x2007: []rune{0x0020},
	0x2008: []rune{0x0020},
	0x2009: []rune{0x0020},
	0x200A: []rune{0x0020},
	0x200B: []rune{0x0020},
	0x202F: []rune{0x0020},
	0x205F: []rune{0x0020},
	0x3000: []rune{0x0020},
}

var saslprep = stringprep.Profile{
	Mappings: []stringprep.Mapping{
		stringprep.TableB1,
		mapNonASCIISpaceToASCIISpace,
	},
	Normalize: true,
	Prohibits: []stringprep.Set{
		stringprep.TableA1,
		stringprep.TableC1_2,
		stringprep.TableC2_1,
		stringprep.TableC2_2,
		stringprep.TableC3,
		stringprep.TableC4,
		stringprep.TableC5,
		stringprep.TableC6,
		stringprep.TableC7,
		stringprep.TableC8,
		stringprep.TableC9,
	},
	CheckBiDi: true,
}

// Prepare converts a string according the SASLprep rules defined in RFC-4013.
//
// Because the stringprep distinction between query and stored strings was
// intended for compatibility across profile versions, but SASLprep was never
// updated and is now deprecated, this function only operates in stored
// strings mode, prohibiting unassigned code points.
func Prepare(s string) (string, error) {
	return saslprep.Prepare(s)
}
