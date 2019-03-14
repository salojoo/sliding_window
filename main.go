package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var LINE_CHANGE = "\r\n"

func main() {
	inputFilename := flag.String("input", "", "Input file")
	outputFilename := flag.String("output", "", "Output file")
	windowSize := flag.Int("window", 5, "Sliding window targetSize")
	algorithm := flag.String("algorithm", "sort", "Algorithm [sort|optimized|bucket]")
	bucketMin := flag.Int("min", 100, "Minimum high resolution value for bucket sort")
	bucketMax := flag.Int("max", 200, "Maximum high resolution value for bucket sort")
	flag.Parse()

	if len(*inputFilename) < 1 || len(*outputFilename) < 1 {
		flag.PrintDefaults()
		return
	}

	fmt.Println("Processing file", *inputFilename)
	fmt.Println("Window targetSize", *windowSize)

	inputFile, err := os.Open(*inputFilename)
	if err != nil {
		fmt.Println("Error opening file", *inputFilename, err)
		return
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)

	outputFile, err := os.Create(*outputFilename)
	if err != nil {
		fmt.Println("Error creating file", *outputFilename, err)
		return
	}
	defer outputFile.Close()

	var sw SlidingWindow
	if *algorithm == "sort" {
		sw = &SlidingWindowSort{targetSize: *windowSize}
	} else if *algorithm == "optimized" {
		sw = &SlidingWindowOptimized{targetSize: *windowSize}
	} else if *algorithm == "bucket" {
		sw = NewSlidingWindowBucket(*windowSize, *bucketMin, *bucketMax)
	}

	writer := bufio.NewWriter(outputFile)

	var lines int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) < 1 {
			continue
		}

		value, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Println("Could not parse integer from line:", line)
			continue
		}

		sw.AddDelay(int(value))

		if _, err := writer.WriteString(strconv.Itoa(sw.GetMedian())); err != nil {
			fmt.Println("Error writing to file", err)
		}
		if _, err := writer.WriteString(LINE_CHANGE); err != nil {
			fmt.Println("Error writing to file", err)
		}

		lines += 1
	}

	if err := writer.Flush(); err != nil {
		fmt.Println("Error flushing file", err)
	}

	fmt.Println("Processed", lines, "lines")
}
