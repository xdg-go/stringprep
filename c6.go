// Copyright 2018 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package stringprep

var tableC6 = Table{
	RuneRange{0xFFF9, 0xFFF9}, // INTERLINEAR ANNOTATION ANCHOR
	RuneRange{0xFFFA, 0xFFFA}, // INTERLINEAR ANNOTATION SEPARATOR
	RuneRange{0xFFFB, 0xFFFB}, // INTERLINEAR ANNOTATION TERMINATOR
	RuneRange{0xFFFC, 0xFFFC}, // OBJECT REPLACEMENT CHARACTER
	RuneRange{0xFFFD, 0xFFFD}, // REPLACEMENT CHARACTER
}

// TableC6 represents RFC-3454 Table C.6.
var TableC6 Table = tableC6
