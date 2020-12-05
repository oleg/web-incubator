#!/bin/bash

#given
make ex-1.1-hello

#when
./ex-1.1-hello >output

#then
actual=$(<output)
expected="hello, world"
if [ "$actual" != "$expected" ]; then
  echo "expected '$expected' but got '$actual'"
  exit 1
fi
