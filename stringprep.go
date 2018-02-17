// Copyright 2018 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

// Package stringprep provides data tables and algorithms for RFC-3454,
// including errata (as of 2018-02).
package stringprep

import "sort"

// RuneRange represents a close-ended range of runes: [N,M].  For a range
// consisting of a single rune, N and M will be equal.
type RuneRange [2]rune

// Contains returns true if a rune is within the bounds of the RuneRange.
func (rr RuneRange) Contains(r rune) bool {
	return rr[0] <= r && r <= rr[1]
}

func (rr RuneRange) isAbove(r rune) bool {
	return r <= rr[0]
}

// Set represents a stringprep data table used to identify runes of a
// particular type.
type Set []RuneRange

// Contains returns true if a rune is within any of the RuneRanges in the
// Set.
func (t Set) Contains(r rune) bool {
	i := sort.Search(len(t), func(i int) bool { return t[i].Contains(r) || t[i].isAbove(r) })
	if i < len(t) && t[i].Contains(r) {
		return true
	}
	return false
}

// Mapping represents a stringprep mapping, from a single rune to zero or more
// runes.
type Mapping map[rune][]rune

// Map maps a rune to a (possibly empty) rune slice via a stringprep Mapping.
// The ok return value is false if the rune was not found.
func (m Mapping) Map(r rune) (replacement []rune, ok bool) {
	rs, ok := m[r]
	if !ok {
		return nil, false
	}
	return rs, true
}
