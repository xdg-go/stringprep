// Copyright 2018 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package stringprep

var tableC3 = Table{
	RuneRange{0xE000, 0xF8FF},     // [PRIVATE USE, PLANE 0]
	RuneRange{0xF0000, 0xFFFFD},   // [PRIVATE USE, PLANE 15]
	RuneRange{0x100000, 0x10FFFD}, // [PRIVATE USE, PLANE 16]
}

// TableC3 represents RFC-3454 Table C.3.
var TableC3 Table = tableC3
