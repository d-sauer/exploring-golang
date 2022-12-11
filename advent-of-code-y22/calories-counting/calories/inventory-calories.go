package calories

import (
	"sort"
)

type InventoryIndex struct {
	Index         []Elf
	TotalCalories int
}

type Elf struct {
	Identifier    int
	StockCalories int
}

func (ic *InventoryIndex) Process(identifier int, totalCalories int) {
	elf := Elf{Identifier: identifier, StockCalories: totalCalories}

	ic.enter(elf)
}

func (ic *InventoryIndex) enter(elf Elf) {
	ic.TotalCalories += elf.StockCalories

	i := sort.Search(len(ic.Index), func(i int) bool { return ic.Index[i].StockCalories <= elf.StockCalories })

	ic.Index = append(ic.Index, Elf{})
	// shift right for one place, from index 'i'
	copy(ic.Index[i+1:], ic.Index[i:])
	ic.Index[i] = elf
}

func (ic InventoryIndex) MaxCalories() int {
	return ic.Index[0].StockCalories
}

func (ic InventoryIndex) MinCalories() int {
	return ic.Index[len(ic.Index)-1].StockCalories
}

func (ic InventoryIndex) SumTopCarriers(limit int) int {
	sum := 0
	for i := 0; i < limit; i++ {
		sum += ic.Index[i].StockCalories
	}

	return sum
}
