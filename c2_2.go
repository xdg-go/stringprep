// Copyright 2018 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package stringprep

var tableC2_2 = Set{
	RuneRange{0x0080, 0x009F},   // [CONTROL CHARACTERS]
	RuneRange{0x06DD, 0x06DD},   // ARABIC END OF AYAH
	RuneRange{0x070F, 0x070F},   // SYRIAC ABBREVIATION MARK
	RuneRange{0x180E, 0x180E},   // MONGOLIAN VOWEL SEPARATOR
	RuneRange{0x200C, 0x200C},   // ZERO WIDTH NON-JOINER
	RuneRange{0x200D, 0x200D},   // ZERO WIDTH JOINER
	RuneRange{0x2028, 0x2028},   // LINE SEPARATOR
	RuneRange{0x2029, 0x2029},   // PARAGRAPH SEPARATOR
	RuneRange{0x2060, 0x2060},   // WORD JOINER
	RuneRange{0x2061, 0x2061},   // FUNCTION APPLICATION
	RuneRange{0x2062, 0x2062},   // INVISIBLE TIMES
	RuneRange{0x2063, 0x2063},   // INVISIBLE SEPARATOR
	RuneRange{0x206A, 0x206F},   // [CONTROL CHARACTERS]
	RuneRange{0xFEFF, 0xFEFF},   // ZERO WIDTH NO-BREAK SPACE
	RuneRange{0xFFF9, 0xFFFC},   // [CONTROL CHARACTERS]
	RuneRange{0x1D173, 0x1D17A}, // [MUSICAL CONTROL CHARACTERS]
}

// TableC2_2 represents RFC-3454 Table C.2.2.
var TableC2_2 Set = tableC2_2
