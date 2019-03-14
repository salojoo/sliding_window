package main

import "sort"

type SlidingWindowOptimized struct {
	targetSize  int
	items       []int
	sortedItems []int
}

func (sw *SlidingWindowOptimized) AddDelay(delay int) {
	sw.items = append(sw.items, delay)
	sw.sortedItems = append(sw.sortedItems, delay)

	if len(sw.items) > sw.targetSize {
		valueToRemove := sw.items[0]
		for i, v := range sw.sortedItems {
			if v == valueToRemove {
				sw.sortedItems = append(sw.sortedItems[:i], sw.sortedItems[i+1:]...)
				break
			}
		}
		sw.items = sw.items[1:]
	}

	sort.Ints(sw.sortedItems)
}

func (sw *SlidingWindowOptimized) GetMedian() int {
	currentSize := len(sw.items)

	if currentSize == 1 {
		return -1
	}

	if currentSize%2 == 1 {
		return sw.sortedItems[(currentSize+1)/2-1]
	} else {
		return (sw.sortedItems[currentSize/2-1] + sw.sortedItems[currentSize/2+1-1]) / 2
	}
}
