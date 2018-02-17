#!/usr/bin/env perl
use v5.10;
use strict;
use warnings;
use utf8;
use open qw/:std :utf8/;
use Path::Tiny;

my $file = shift(@ARGV) or die "Usage: $0 <file>";
my $path = path($file);
my $base = $path->basename(".txt");
my $table = $base;
$table =~ s/_//; # first only
$base =~ s/_/./g; # back to dots for docs

my @output = <<"HERE";
// Copyright 2018 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package stringprep

import "sort"

var table$table = Table{
HERE

for my $line ( $path->lines( { chomp => 1 } ) ) {
    if (
        $line =~ m{
        ^
        (\w+)               # start rune
        (?:-(\w+))?      # end rune
        (?:;\s+(.*))?   # comment
        $
        }x
      )
    {
        my ( $start, $end, $comment ) = map { $_ // "" } ( $1, $2, $3 );
        $end = $start unless length $end;
        my $code = "\tRuneRange{0x$start, 0x$end},";
        $code .= " // $comment" if length $comment;
        push @output, $code . "\n";
    }
    else {
        die "match failed on '$line'\n";
    }
}

push @output, <<"here";
}

// Table$table represents RFC-3454 Table $base.
var Table$table Table = table$table

// Is$table returns true if a rune is in stringprep table $base.
func Is$table(r rune) bool {
    i := sort.Search( len(table$table), func(i int) bool { return table$table\[i].Contains(r) || table$table\[i][0] >= r })
    if i < len(table$table) && table$table\[i].Contains(r) {
        return true
    }
    return false
}
here

path(lc("$table.go"))->spew(@output);
