package main

import "testing"

func TestSlidingWindow_AddDelay(t *testing.T) {
	sw := SlidingWindow{targetSize: 2}
	sw.AddDelay(1)
	sw.AddDelay(2)
	sw.AddDelay(3)

}

func TestSlidingWindow_GetMedian(t *testing.T) {
	sw := SlidingWindow{targetSize: 3}

	sw.AddDelay(100)
	if sw.GetMedian() != -1 {
		t.Errorf("expected %d, got %d", -1, sw.GetMedian())
	}

	sw.AddDelay(102)
	if sw.GetMedian() != 101 {
		t.Errorf("expected %d, got %d", 101, sw.GetMedian())
	}

	sw.AddDelay(101)
	if sw.GetMedian() != 101 {
		t.Errorf("expected %d, got %d", 101, sw.GetMedian())
	}

	sw.AddDelay(110)
	if sw.GetMedian() != 102 {
		t.Errorf("expected %d, got %d", 102, sw.GetMedian())
	}

	sw.AddDelay(120)
	if sw.GetMedian() != 110 {
		t.Errorf("expected %d, got %d", 101, sw.GetMedian())
	}

	sw.AddDelay(115)
	if sw.GetMedian() != 115 {
		t.Errorf("expected %d, got %d", 115, sw.GetMedian())
	}
}
