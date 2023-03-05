package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	var fileProperty = flag.String("file", "", "input file")

	flag.Parse()

	fmt.Printf("Opening file: %s\n", *fileProperty)
	file, err := os.Open(*fileProperty)

	if err != nil {
		fmt.Printf("can't open file %s: %v", *fileProperty, err)
		os.Exit(1)
	}
	defer file.Close()

	count, err := CountCollisionPairs(file)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Assignment collision count: %d\n", count)
}

func CountCollisionPairs(file fs.File) (int, error) {
	sc := bufio.NewScanner(file)

	var line, collideCount int

	for sc.Scan() {
		line++

		content := sc.Text()

		r1 := Range{}
		r2 := Range{}
		p, err := fmt.Sscanf(content, "%d-%d,%d-%d", &r1.Start, &r1.End, &r2.Start, &r2.End)

		if err != nil || p != 4 {
			return -1, fmt.Errorf("error parsin line: %d %w", line, err)
		}

		collide := IsColliding(r1, r2)
		if collide {
			collideCount++
		}
	}

	return collideCount, nil
}
