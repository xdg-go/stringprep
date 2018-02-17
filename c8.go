// Copyright 2018 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package stringprep

var tableC8 = Table{
	RuneRange{0x0340, 0x0340}, // COMBINING GRAVE TONE MARK
	RuneRange{0x0341, 0x0341}, // COMBINING ACUTE TONE MARK
	RuneRange{0x200E, 0x200E}, // LEFT-TO-RIGHT MARK
	RuneRange{0x200F, 0x200F}, // RIGHT-TO-LEFT MARK
	RuneRange{0x202A, 0x202A}, // LEFT-TO-RIGHT EMBEDDING
	RuneRange{0x202B, 0x202B}, // RIGHT-TO-LEFT EMBEDDING
	RuneRange{0x202C, 0x202C}, // POP DIRECTIONAL FORMATTING
	RuneRange{0x202D, 0x202D}, // LEFT-TO-RIGHT OVERRIDE
	RuneRange{0x202E, 0x202E}, // RIGHT-TO-LEFT OVERRIDE
	RuneRange{0x206A, 0x206A}, // INHIBIT SYMMETRIC SWAPPING
	RuneRange{0x206B, 0x206B}, // ACTIVATE SYMMETRIC SWAPPING
	RuneRange{0x206C, 0x206C}, // INHIBIT ARABIC FORM SHAPING
	RuneRange{0x206D, 0x206D}, // ACTIVATE ARABIC FORM SHAPING
	RuneRange{0x206E, 0x206E}, // NATIONAL DIGIT SHAPES
	RuneRange{0x206F, 0x206F}, // NOMINAL DIGIT SHAPES
}

// TableC8 represents RFC-3454 Table C.8.
var TableC8 Table = tableC8
