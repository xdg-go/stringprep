// Copyright 2018 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package stringprep

var tableC2_1 = Set{
	RuneRange{0x0000, 0x001F}, // [CONTROL CHARACTERS]
	RuneRange{0x007F, 0x007F}, // DELETE
}

// TableC2_1 represents RFC-3454 Table C.2.1.
var TableC2_1 Set = tableC2_1
