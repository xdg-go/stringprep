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

var table$table = Mapping{
HERE

for my $line ( $path->lines( { chomp => 1 } ) ) {
    my ( $in, $out, $comment ) = split /; +/, $line;
    my @replace = split / /, $out;
    my $out_range = "[]rune{" . join( ",", map { "0x$_" } @replace ) . "}";
    my $code = "\t0x$in:\t$out_range,";
    $code .= "\t// $comment" if length $comment;
    push @output, $code . "\n";
}

push @output, <<"HERE";
}

// Table$table represents RFC-3454 Table $base.
var Table$table Mapping = table$table
HERE

path( lc("$table.go") )->spew(@output);
