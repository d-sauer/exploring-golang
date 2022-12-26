package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"testing/fstest"
)

func Test_itemPriority(t *testing.T) {
	priority := itemPriority('a')
	assert.Equal(t, 1, priority)

	priority = itemPriority('z')
	assert.Equal(t, 26, priority)

	priority = itemPriority('A')
	assert.Equal(t, 27, priority)

	priority = itemPriority('Z')
	assert.Equal(t, 52, priority)

	priority = itemPriority('8')
	assert.Equal(t, 0, priority)
}

func Test_loadInput(t *testing.T) {
	tf := fstest.MapFS{"content.txt": {Data: []byte(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`)}}

	file, err := tf.Open("content.txt")
	if err != nil {
		t.Error(err)
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

func Test_findDuplicateItems(t *testing.T) {
	duplicates, _ := findDuplicatedItems("vJrwpWtwJgWrhcsFMMfFFhFp")
	assert.Contains(t, duplicates, "p")

	duplicates, _ = findDuplicatedItems("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL")
	assert.Contains(t, duplicates, "L")
}

func Test_findGroupBadge(t *testing.T) {
	group := []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}

	item := findGroupBadge(&group)
	assert.Contains(t, item, "r")
}
