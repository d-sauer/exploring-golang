package calories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInventoryIndex_enter(t *testing.T) {
	e1 := Elf{
		Identifier:    1,
		StockCalories: 3,
	}

	e2 := Elf{
		Identifier:    3,
		StockCalories: 2,
	}

	e3 := Elf{
		Identifier:    3,
		StockCalories: 1,
	}

	e4 := Elf{
		Identifier:    4,
		StockCalories: 4,
	}

	ic := &InventoryIndex{}
	ic.enter(e1)
	ic.enter(e2)
	ic.enter(e3)
	ic.enter(e4)

	assert.Equal(t, 4, ic.MaxCalories(), "Max amount of calories is not as expected")
	assert.Equal(t, 1, ic.MinCalories(), "Min amount of calories is not as expected")
	assert.Equal(t, 9, ic.SumTopCarriers(3), "Sum of top carriers is not as expected")
}
