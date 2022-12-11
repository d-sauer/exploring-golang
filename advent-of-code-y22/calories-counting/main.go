package main

import (
	"flag"
	"fmt"
	calories "github.com/d-sauer/exploring-go/advent-of-code-y22/calories-counting/calories"
	"os"
)

func main() {
	var fileProperty = flag.String("file", "", "Inventory file")
	flag.Parse()

	fmt.Printf("Opening inventory file: %s\n", *fileProperty)

	file, err := os.Open(*fileProperty)
	if err != nil {
		fmt.Errorf("can't open file %s: %w", *fileProperty, err)
	}
	defer file.Close()

	ic := calories.InventoryIndex{}
	err = calories.LoadInventory(file, &ic)
	if err != nil {
		fmt.Errorf("can't load inverntory %s: %w", *fileProperty, err)
		return
	}

	var ics InventoryStatistics = ic // experimenting with Interfaces
	max := ics.MaxCalories()
	min := ics.MinCalories()
	topThreeSum := ics.SumTopCarriers(3)

	fmt.Printf("Max calories: %d, Min calories: %d, Sum calories of top three elves: %d\n", max, min, topThreeSum)
}
