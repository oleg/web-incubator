#!/bin/sh -ex

rm db
./ex17-d db c
./ex17-d db l
./ex17-d db s 1 "Oleg Pro" "oleg@example.com"
./ex17-d db s 2 "John Doe" "john.doe@example.com"
./ex17-d db l
./ex17-d db g 1
./ex17-d db d 1
./ex17-d db l
