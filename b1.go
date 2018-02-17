// Copyright 2018 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package stringprep

var tableB1 = Mapping{
	0x00AD: []rune{}, // Map to nothing
	0x034F: []rune{}, // Map to nothing
	0x180B: []rune{}, // Map to nothing
	0x180C: []rune{}, // Map to nothing
	0x180D: []rune{}, // Map to nothing
	0x200B: []rune{}, // Map to nothing
	0x200C: []rune{}, // Map to nothing
	0x200D: []rune{}, // Map to nothing
	0x2060: []rune{}, // Map to nothing
	0xFE00: []rune{}, // Map to nothing
	0xFE01: []rune{}, // Map to nothing
	0xFE02: []rune{}, // Map to nothing
	0xFE03: []rune{}, // Map to nothing
	0xFE04: []rune{}, // Map to nothing
	0xFE05: []rune{}, // Map to nothing
	0xFE06: []rune{}, // Map to nothing
	0xFE07: []rune{}, // Map to nothing
	0xFE08: []rune{}, // Map to nothing
	0xFE09: []rune{}, // Map to nothing
	0xFE0A: []rune{}, // Map to nothing
	0xFE0B: []rune{}, // Map to nothing
	0xFE0C: []rune{}, // Map to nothing
	0xFE0D: []rune{}, // Map to nothing
	0xFE0E: []rune{}, // Map to nothing
	0xFE0F: []rune{}, // Map to nothing
	0xFEFF: []rune{}, // Map to nothing
}

// TableB1 represents RFC-3454 Table B.1.
var TableB1 Mapping = tableB1
