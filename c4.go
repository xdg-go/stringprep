// Copyright 2018 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package stringprep

var tableC4 = Table{
	RuneRange{0xFDD0, 0xFDEF},     // [NONCHARACTER CODE POINTS]
	RuneRange{0xFFFE, 0xFFFF},     // [NONCHARACTER CODE POINTS]
	RuneRange{0x1FFFE, 0x1FFFF},   // [NONCHARACTER CODE POINTS]
	RuneRange{0x2FFFE, 0x2FFFF},   // [NONCHARACTER CODE POINTS]
	RuneRange{0x3FFFE, 0x3FFFF},   // [NONCHARACTER CODE POINTS]
	RuneRange{0x4FFFE, 0x4FFFF},   // [NONCHARACTER CODE POINTS]
	RuneRange{0x5FFFE, 0x5FFFF},   // [NONCHARACTER CODE POINTS]
	RuneRange{0x6FFFE, 0x6FFFF},   // [NONCHARACTER CODE POINTS]
	RuneRange{0x7FFFE, 0x7FFFF},   // [NONCHARACTER CODE POINTS]
	RuneRange{0x8FFFE, 0x8FFFF},   // [NONCHARACTER CODE POINTS]
	RuneRange{0x9FFFE, 0x9FFFF},   // [NONCHARACTER CODE POINTS]
	RuneRange{0xAFFFE, 0xAFFFF},   // [NONCHARACTER CODE POINTS]
	RuneRange{0xBFFFE, 0xBFFFF},   // [NONCHARACTER CODE POINTS]
	RuneRange{0xCFFFE, 0xCFFFF},   // [NONCHARACTER CODE POINTS]
	RuneRange{0xDFFFE, 0xDFFFF},   // [NONCHARACTER CODE POINTS]
	RuneRange{0xEFFFE, 0xEFFFF},   // [NONCHARACTER CODE POINTS]
	RuneRange{0xFFFFE, 0xFFFFF},   // [NONCHARACTER CODE POINTS]
	RuneRange{0x10FFFE, 0x10FFFF}, // [NONCHARACTER CODE POINTS]
}

// TableC4 represents RFC-3454 Table C.4.
var TableC4 Table = tableC4
