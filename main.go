package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lineChange = "\r\n"

type Config struct {
	inputFilename  *string
	outputFilename *string
	windowSize     *int
	algorithm      *string
	bucketMin      *int
	bucketMax      *int
}

func main() {
	var config Config
	if err := parseParameters(&config); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Processing file", *config.inputFilename)
	fmt.Println("Window targetSize", *config.windowSize)

	if err := process(config); err != nil {
		fmt.Println("Error in processing", err)
	}
}

func parseParameters(config *Config) error {
	config.inputFilename = flag.String("input", "", "Input file")
	config.outputFilename = flag.String("output", "", "Output file")
	config.windowSize = flag.Int("window", 5, "Sliding window targetSize")
	config.algorithm = flag.String("algorithm", "naive", "Algorithm [naive|optimized|bucket]")
	config.bucketMin = flag.Int("min", 100, "Minimum high resolution value for bucket algorithm")
	config.bucketMax = flag.Int("max", 200, "Maximum high resolution value for bucket algorithm")
	flag.Parse()

	if len(*config.inputFilename) < 1 || len(*config.outputFilename) < 1 {
		flag.PrintDefaults()
		return errors.New("invalid config")
	}

	return nil
}

func createSlidingWindow(config Config) SlidingWindow {
	if *config.algorithm == "naive" {
		return &SlidingWindowNaive{targetSize: *config.windowSize}
	} else if *config.algorithm == "optimized" {
		return &SlidingWindowOptimized{targetSize: *config.windowSize}
	} else if *config.algorithm == "bucket" {
		return NewSlidingWindowBucket(*config.windowSize, *config.bucketMin, *config.bucketMax)
	}
	return nil
}

func process(config Config) (err error) {
	inputFile, err := os.Open(*config.inputFilename)
	if err != nil {
		fmt.Println("Error opening file", *config.inputFilename, err)
		return
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)

	outputFile, err := os.Create(*config.outputFilename)
	if err != nil {
		fmt.Println("Error creating file", *config.outputFilename, err)
		return
	}
	defer outputFile.Close()
	writer := bufio.NewWriter(outputFile)

	sw := createSlidingWindow(config)

	var lines int
	for scanner.Scan() {
		var value int
		if value, err = readLine(scanner); err != nil {
			fmt.Println("Skipping line:", err)
			continue
		}

		sw.AddDelay(value)

		if err = writeLine(sw.GetMedian(), writer); err != nil {
			fmt.Println("Could not write to file:", err)
			return
		}

		lines += 1
	}

	if err := writer.Flush(); err != nil {
		fmt.Println("Error flushing file", err)
	}

	fmt.Println("Processed", lines, "lines")

	return nil
}

func readLine(scanner *bufio.Scanner) (int, error) {
	line := strings.TrimSpace(scanner.Text())
	if len(line) < 1 {
		return 0, errors.New("short line")
	}

	value, err := strconv.ParseInt(line, 10, 64)
	return int(value), err
}

func writeLine(value int, writer *bufio.Writer) (err error) {
	if _, err = writer.WriteString(strconv.Itoa(value)); err != nil {
		return
	}
	_, err = writer.WriteString(lineChange)

	return
}
