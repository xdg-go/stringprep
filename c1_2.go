// Copyright 2018 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package stringprep

var tableC1_2 = Table{
	RuneRange{0x00A0, 0x00A0}, // NO-BREAK SPACE
	RuneRange{0x1680, 0x1680}, // OGHAM SPACE MARK
	RuneRange{0x2000, 0x2000}, // EN QUAD
	RuneRange{0x2001, 0x2001}, // EM QUAD
	RuneRange{0x2002, 0x2002}, // EN SPACE
	RuneRange{0x2003, 0x2003}, // EM SPACE
	RuneRange{0x2004, 0x2004}, // THREE-PER-EM SPACE
	RuneRange{0x2005, 0x2005}, // FOUR-PER-EM SPACE
	RuneRange{0x2006, 0x2006}, // SIX-PER-EM SPACE
	RuneRange{0x2007, 0x2007}, // FIGURE SPACE
	RuneRange{0x2008, 0x2008}, // PUNCTUATION SPACE
	RuneRange{0x2009, 0x2009}, // THIN SPACE
	RuneRange{0x200A, 0x200A}, // HAIR SPACE
	RuneRange{0x200B, 0x200B}, // ZERO WIDTH SPACE
	RuneRange{0x202F, 0x202F}, // NARROW NO-BREAK SPACE
	RuneRange{0x205F, 0x205F}, // MEDIUM MATHEMATICAL SPACE
	RuneRange{0x3000, 0x3000}, // IDEOGRAPHIC SPACE
}

// TableC1_2 represents RFC-3454 Table C.1.2.
var TableC1_2 Table = tableC1_2
