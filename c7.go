// Copyright 2018 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package stringprep

var tableC7 = Set{
	RuneRange{0x2FF0, 0x2FFB}, // [IDEOGRAPHIC DESCRIPTION CHARACTERS]
}

// TableC7 represents RFC-3454 Table C.7.
var TableC7 Set = tableC7
