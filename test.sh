#!/bin/bash

# rm -f calendar.json

t="2023-12-04 22:02"

./calstatus --time 8M  --text "e"
./calstatus --time 2M  --text "a"
./calstatus --time 6M  --text "c"
./calstatus --time 4M  --text "b"

./calstatus --at "$t"  --text "d"

./calstatus --time 12M --text "f"
./calstatus --time 16M --text "h"
./calstatus --time 14M --text "g"


# jq -r 'sort_by(.Time)' calendar.json
jq -r . calendar.json
