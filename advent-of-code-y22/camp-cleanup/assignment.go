package main

type Range struct {
	Start int
	End   int
}

// AssignemntCollide detect if RangeA is withing RangeB or otherwise around
func AssignemntCollide(RangeA Range, RangeB Range) bool {
	if RangeB.Start >= RangeA.Start && RangeB.End <= RangeA.End {
		return true
	} else if RangeA.Start >= RangeB.Start && RangeA.End <= RangeB.End {
		return true
	}

	return false
}
