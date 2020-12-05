#!/bin/bash
set -e

#given
make ex-1.3-fahr-cel

#when
./ex-1.3-fahr-cel >ex-1.3-fahr-cel.actual

#then
diff ex-1.3-fahr-cel.actual ex-1.3-fahr-cel.expected
