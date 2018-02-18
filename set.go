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
func (s Set) Contains(r rune) bool {
	i := sort.Search(len(s), func(i int) bool { return s[i].Contains(r) || s[i].isAbove(r) })
	if i < len(s) && s[i].Contains(r) {
		return true
	}
	return false
}

// Check for prohibited characters from table C.8
func hasBiDiProhibitedRune(s string) bool {
	for _, r := range s {
		if TableC8.Contains(r) {
			return true
		}
	}
	return false
}

// Check for RandALCat characters from table D.1
func hasBiDiRandALCat(s string) bool {
	for _, r := range s {
		if TableD1.Contains(r) {
			return true
		}
	}
	return false
}

// Check for LCat characters from table D.2
func hasBiDiLCat(s string) bool {
	for _, r := range s {
		if TableD2.Contains(r) {
			return true
		}
	}
	return false
}

// Check first and last characters are in table D.1; requires non-empty string
func hasFirstAndLastRandALCat(s string) bool {
	rs := []rune(s)
	return TableD1.Contains(rs[0]) && TableD1.Contains(rs[len(rs)-1])
}

// Check that BiDi rules are satisfied ; let empty string pass this rule
func passesBiDiRules(s string) bool {
	if len(s) == 0 {
		return true
	}
	if hasBiDiProhibitedRune(s) {
		return false
	}
	if hasBiDiRandALCat(s) {
		if hasBiDiLCat(s) {
			return false
		}
		if !hasFirstAndLastRandALCat(s) {
			return false
		}
	}
	return true
}
