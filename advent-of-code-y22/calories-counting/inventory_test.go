package main_test

import (
	"fmt"
	calories "github.com/d-sauer/exploring-go/advent-of-code-y22/calories-counting/calories"
	"testing"
	"testing/fstest"
)

func TestInput(t *testing.T) {
	tf := fstest.MapFS{"inventory.txt": {Data: []byte(`1
    1
         
    2
    2

    3
    3
`)}}

	ic := calories.InventoryIndex{}

	file, err := tf.Open("inventory.txt")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	calories.LoadInventory(file, &ic)

	max := ic.MaxCalories()
	min := ic.MinCalories()
	topThreeSum := ic.SumTopCarriers(3)

	fmt.Printf("Max calories: %d, Min calories: %d, Sum calories of top three elves: %d", max, min, topThreeSum)
}
