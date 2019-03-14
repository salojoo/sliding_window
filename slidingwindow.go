package main

type SlidingWindow interface {
	AddDelay(int)
	GetMedian() int
}
