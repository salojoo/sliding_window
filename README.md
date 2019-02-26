# SlidingWindow

## Running
Run from command line with -help for instructions

## Test data
Test data available in data directory

## Optimizations
Current implementation is very inefficient and could be improved at least in two different ways

### Parallel
Input files could be split so sliding windows could run in parallel. This would give performance boost of 2x, 4x, 8x, or the number of cores available

### Optimizing the sort algorithm
In the worst case the sort algorithm needs to sort 10000 items for each window. A more efficient flow would be to keep a sorted window in memory. Then on AddDelay():
- remove oldest value 
- add new value
- sort again

As 9999 items of 10000 are already sorted, this would only require finding the place for one item on each iteration of AddDelay()

Optimizing the sort algorithm should give more performance than parallel processing if window sizes are big.
