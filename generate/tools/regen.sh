#!/bin/bash

find ./$1 -type f -name "*.gen.go" -exec dirname {} \; | sort -u | xargs -I{} go generate {}