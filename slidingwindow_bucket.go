package main

type SlidingWindowBucket struct {
	targetSize int
	items      []int

	// value range for "high-resolution" bucket counter
	min int
	max int

	bucket    []int // items for "high-resolution" range of values
	minBucket int   // counting items of lesser value than min
	maxBucket int   // counting items of higher value than max
}

func NewSlidingWindowBucket(targetSize int, min int, max int) (sw *SlidingWindowBucket) {
	sw = &SlidingWindowBucket{targetSize: targetSize, min: min, max: max}
	sw.bucket = make([]int, max-min+1)
	return
}

func (sw *SlidingWindowBucket) AddDelay(delay int) {
	sw.items = append(sw.items, delay)

	// constant insertion time
	if delay < sw.min {
		sw.minBucket += 1
	} else if delay > sw.max {
		sw.maxBucket += 1
	} else {
		sw.bucket[delay-sw.min] += 1
	}

	if len(sw.items) > sw.targetSize {
		valueToRemove := sw.items[0]
		sw.items = sw.items[1:]

		// constant removal time
		if valueToRemove < sw.min {
			sw.minBucket -= 1
		} else if valueToRemove > sw.max {
			sw.maxBucket -= 1
		} else {
			sw.bucket[valueToRemove-sw.min] -= 1
		}
	}
}

func (sw *SlidingWindowBucket) GetMedian() int {
	currentSize := len(sw.items)

	if currentSize == 1 {
		return -1
	}

	if currentSize%2 == 1 {
		return sw.findValueFromBucket((currentSize+1)/2 - 1)
		//return sortedItems[(currentSize+1)/2-1]
	} else {
		value1 := sw.findValueFromBucket(currentSize/2 - 1)
		value2 := sw.findValueFromBucket(currentSize/2 + 1 - 1)
		return (value1 + value2) / 2
		//return (sortedItems[currentSize/2-1] + sortedItems[currentSize/2+1-1]) / 2
	}
}

func (sw *SlidingWindowBucket) findValueFromBucket(index int) (bucketValue int) {
	search := sw.minBucket

	if index < search {
		panic("Bucket is too small")
	}

	for i, count := range sw.bucket {
		search += count
		if search > index {
			bucketValue = i + sw.min
			return
		}
	}

	panic("Bucket is too small")
}
