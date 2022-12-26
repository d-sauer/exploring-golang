package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var regex = regexp.MustCompile("(?P<S1>\\d+)-(?P<E1>\\d+),(?P<S2>\\d+)-(?P<E2>\\d+)")

func main() {
	var fileProperty = flag.String("file", "", "input file")
	flag.Parse()

	fmt.Printf("Opening file: %s\n", *fileProperty)

	file, err := os.Open(*fileProperty)
	if err != nil {
		fmt.Errorf("can't open file %s: %w", *fileProperty, err)
	}
	defer file.Close()

	count, err := CountCollisionPairs(file)
	if err != nil {
		fmt.Errorf("can't count collision pairs due to %w'", err)
	}
	fmt.Printf("Assignment collision count: %d\n", count)
}

func CountCollisionPairs(file fs.File) (int, error) {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var line = 0
	var collideCount = 0
	for fileScanner.Scan() {
		line++
		content := strings.TrimSpace(fileScanner.Text())

		r1, r2, err := parseContent(content)
		if err != nil {
			return -1, fmt.Errorf("error parsin line: %d, %w\n", line, err)
		}

		collide := AssignemntCollide(r1, r2)
		if collide {
			collideCount++
		}
	}

	return collideCount, nil
}

func parseContent(content string) (Range, Range, error) {
	match := regex.FindStringSubmatch(content)
	if len(match) == 0 {
		return Range{}, Range{}, fmt.Errorf("no matches")
	}

	s1, _ := strconv.Atoi(match[regex.SubexpIndex("S1")])
	e1, _ := strconv.Atoi(match[regex.SubexpIndex("E1")])
	r1 := Range{Start: s1, End: e1}

	s2, _ := strconv.Atoi(match[regex.SubexpIndex("S2")])
	e2, _ := strconv.Atoi(match[regex.SubexpIndex("E2")])
	r2 := Range{Start: s2, End: e2}

	return r1, r2, nil
}
