package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var fileProperty = flag.String("file", "", "rucksack content")
	flag.Parse()

	fmt.Printf("Opening file: %s\n", *fileProperty)

	file, err := os.Open(*fileProperty)
	if err != nil {
		fmt.Errorf("can't open file %s: %w", *fileProperty, err)
	}
	defer file.Close()

	duplicates, badges, _ := loadInput(file)

	var sumPriorities int
	for _, duplicate := range duplicates {
		fmt.Println(duplicate)
		sumPriorities += itemPriority(duplicate[0])
	}
	fmt.Println(sumPriorities)

	var sumBadgePriorities int
	for _, badge := range badges {
		fmt.Println(badge)
		sumBadgePriorities += itemPriority(badge[0])
	}
	fmt.Println(sumBadgePriorities)
}
