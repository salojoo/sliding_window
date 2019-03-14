package main

import "sort"

type SlidingWindowSort struct {
	targetSize int
	items      []int
}

func (sw *SlidingWindowSort) AddDelay(delay int) {
	sw.items = append(sw.items, delay)
	if len(sw.items) > sw.targetSize {
		sw.items = sw.items[1:]
	}
}

func (sw *SlidingWindowSort) GetMedian() int {
	currentSize := len(sw.items)

	if currentSize == 1 {
		return -1
	}

	sortedItems := make([]int, len(sw.items))
	copy(sortedItems, sw.items)
	sort.Ints(sortedItems)

	if currentSize%2 == 1 {
		return sortedItems[(currentSize+1)/2-1]
	} else {
		return (sortedItems[currentSize/2-1] + sortedItems[currentSize/2+1-1]) / 2
	}
}
