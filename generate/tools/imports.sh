#!/bin/bash

imports() {
  awk '/^import \(/,/\)/ { if ($1 != "import" && $1 != "(" && $1 != ")") print $1 }' $1 | tr -d '"'
}

for file in $1/*.go; do
    if [ -f "$file" ]; then
        imports "$file"
    fi
done | sort | uniq