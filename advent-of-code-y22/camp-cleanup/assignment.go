package main

type Range struct {
	Start int
	End   int
}

// IsColliding detect if RangeA is within RangeB or otherwise around
func IsColliding(first, second Range) bool {
	if second.Start >= first.Start && second.End <= first.End || first.Start >= second.Start && first.End <= second.End {
		return true
	}

	return false
}
