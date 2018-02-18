[![GoDoc](https://godoc.org/github.com/xdg/stringprep?status.svg)](https://godoc.org/github.com/xdg/stringprep)
[![Build Status](https://travis-ci.org/xdg/stringprep.svg?branch=master)](https://travis-ci.org/xdg/stringprep)

# stringprep – Go implementation of RFC-3454 stringprep

## Synopsis

```
    myProfile := stringprep.Profile{
	Mappings:  []stringprep.Mapping{
            stringprep.TableB1,
            stringprep.TableB2,
        },
	Normalize: true,
	Prohibits: []Set{
            stringprep.TableC1_1,
            stringprep.TableC1_2,
        },
	CheckBiDi: true,
    }

    prepped := myProfile.Prepare(input)

```

## Description

This library provides an implementation of the stringprep algorithm
(RFC-3454), including all data tables, in Go.

The repo includes a SASLprep (RFC-4013) profile built using the stringprep
library. See [SASLprep docs](https://godoc.org/github.com/xdg/stringprep/saslprep)
for usage information.

## Copyright and License

Copyright 2018 by David A. Golden. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License"). You may
obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
