#!/bin/bash

echo "Generated frameworks:"
find ./macos -type f -name "*.gen.go" -exec dirname {} \; | sort -u | wc -l
echo
echo "Classes/protocols:"
find ./macos -type f ! -name "*_test.go" ! -name "*_structs.go" ! -name "aliastypes.gen.go" ! -name "enumtypes.gen.go" ! -name "doc.gen.go" ! -name "protocols.gen.m" -name "*.gen.go" | wc -l
echo
echo "Enums/constants:"
cat ./macos/*/enumtypes.gen.go | grep -v '^\s*//' | grep -v '^\s*$' | grep -v '[\(\)]' | wc -l