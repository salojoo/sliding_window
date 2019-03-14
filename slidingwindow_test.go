package main

import "testing"

func TestSlidingWindowSort_AddDelay(t *testing.T) {
	sw := SlidingWindowSort{targetSize: 2}
	sw.AddDelay(1)
	sw.AddDelay(2)
	sw.AddDelay(3)

}

func TestSlidingWindowSort_GetMedian(t *testing.T) {
	sw := SlidingWindowSort{targetSize: 3}

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

func TestSlidingWindowOptimized_AddDelay(t *testing.T) {
	sw := SlidingWindowOptimized{targetSize: 2}
	sw.AddDelay(1)
	sw.AddDelay(2)
	sw.AddDelay(3)

}

func TestSlidingWindowOptimized_GetMedian(t *testing.T) {
	sw := SlidingWindowOptimized{targetSize: 3}

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

func TestSlidingWindowBucket_AddDelay(t *testing.T) {
	sw := NewSlidingWindowBucket(2, 1, 3)
	sw.AddDelay(1)
	sw.AddDelay(2)
	sw.AddDelay(3)
}

func TestSlidingWindowBucket_GetMedian(t *testing.T) {
	sw := NewSlidingWindowBucket(3, 100, 120)

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
