// Copyright 2018 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package stringprep

var tableC5 = Set{
	RuneRange{0xD800, 0xDFFF}, // [SURROGATE CODES]
}

// TableC5 represents RFC-3454 Table C.5.
var TableC5 Set = tableC5
