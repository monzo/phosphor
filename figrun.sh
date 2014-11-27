#!/bin/sh
set -e

# Pull dependencies and build
go get -v

# Run our binary!
phosphord
