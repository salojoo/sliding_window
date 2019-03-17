# SlidingWindow

## Running
Run from command line with -help for instructions

## Test data
Test data available in data directory

## Benchmarks
Three benchmarks were run and timed on local desktop. Two first use traditional sorting, while the last one uses bucket sort. 

Bucket sort makes the application instantaneous and easily capable of realtime processing. The drawback with bucket sort, is that it needs a predefined range of values as the size of the bucket. It can only be applied if the value range is small enough


### Original naive solution
$ time ./sliding_window -input data/test4.csv -output data/naive4.csv -window 10000 -algorithm naive

Processing file data/test4.csv

Window targetSize 10000

Processed 100000 lines

real	2m44,994s
user	2m44,729s
sys	0m0,321s

### Optimized by using incremental sort instead of recreating the whole window
$ time ./sliding_window -input data/test4.csv -output data/optimized4.csv -window 10000 -algorithm optimized

Processing file data/test4.csv

Window targetSize 10000

Processed 100000 lines

real	1m26,080s
user	1m26,056s
sys	0m0,060s

### Bucket sort
$ time ./sliding_window -input data/test4.csv -output data/bucket4.csv -window 10000 -algorithm bucket -min 200 -max 500

Processing file data/test4.csv

Window targetSize 10000

Processed 100000 lines

real	0m0,058s
user	0m0,057s
sys	0m0,000s



