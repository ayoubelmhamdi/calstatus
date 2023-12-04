#!/bin/bash

# rm -f texts.json

t="2023-12-04 21:54"

./calstatus --time 8M  --add "e"
./calstatus --time 2M  --add "a"
./calstatus --time 6M  --add "c"
./calstatus --time 4M  --add "b"

./calstatus --at "$t"  --add "d"

./calstatus --time 12M --add "f"
./calstatus --time 16M --add "h"
./calstatus --time 14M --add "g"


# jq -r 'sort_by(.Time)' texts.json
jq -r . texts.json
