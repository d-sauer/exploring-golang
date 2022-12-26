package main

import (
	"bufio"
	"fmt"
	"io/fs"
)

type void struct{}

var member void

const (
	groupSize = 3
)

func loadInput(file fs.File) ([]string, []string, error) {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var line = 0
	var allDuplicates []string
	var allBadges []string
	var group = []string{}
	for fileScanner.Scan() {
		line++
		content := fileScanner.Text()

		duplicates, err := findDuplicatedItems(content)
		if err != nil {
			return nil, nil, fmt.Errorf("problem with rucksack number: %d.  %w", line, err)
		}
		allDuplicates = append(allDuplicates, duplicates...)

		g := line % groupSize
		fmt.Println(g)
		group = append(group, content)
		if g == 0 {
			allBadges = append(allBadges, findGroupBadge(&group))
			group = []string{}
		}
	}

	return allDuplicates, allBadges, nil
}

// findGroupBadge finds badge item by looking for item appearing in each group.
func findGroupBadge(group *[]string) string {
	itemCount := map[string]int{}

	for _, g := range *group {
		uniqueItems := make(map[string]void)
		for _, i := range g {
			if _, exists := uniqueItems[string(i)]; !exists {
				uniqueItems[string(i)] = member
				count := itemCount[string(i)] + 1
				itemCount[string(i)] = count
				if count == len(*group) {
					return string(i)
				}
			}
		}
	}

	return ""
}

func findDuplicatedItems(rucksackContent string) ([]string, error) {
	var duplicates []string
	if len(rucksackContent)%2 != 0 {
		return duplicates, fmt.Errorf("rucksack content is not balanced between compartments. Content of rucsack %d items", len(rucksackContent))
	}

	compOne := rucksackContent[:len(rucksackContent)/2]
	compTwo := rucksackContent[len(rucksackContent)/2:]

	for _, c1 := range compOne {
		for _, c2 := range compTwo {
			if c1 == c2 && !contains(duplicates, string(c2)) {
				duplicates = append(duplicates, string(c1))
			}
		}
	}

	return duplicates, nil
}

// itemLetter returns the letter priority as
// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
func itemPriority(itemLetter byte) int {
	if itemLetter >= 97 && itemLetter <= 122 {
		return int(itemLetter - 96)
	}
	if itemLetter >= 65 && itemLetter <= 90 {
		return int(itemLetter - 38)
	}

	return 0
}

func contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
