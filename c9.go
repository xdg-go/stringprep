// Copyright 2018 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package stringprep

var tableC9 = Set{
	RuneRange{0xE0001, 0xE0001}, // LANGUAGE TAG
	RuneRange{0xE0020, 0xE007F}, // [TAGGING CHARACTERS]
}

// TableC9 represents RFC-3454 Table C.9.
var TableC9 Set = tableC9
