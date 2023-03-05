package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignmentCollide(t *testing.T) {
	r1 := Range{Start: 1, End: 5}
	r2 := Range{Start: 2, End: 4}

	collide1 := IsColliding(r1, r2)
	collide2 := IsColliding(r2, r1)

	assert.Equal(t, true, collide1)
	assert.Equal(t, true, collide2)

	r1 = Range{Start: 1, End: 5}
	r2 = Range{Start: 1, End: 4}

	collide1 = IsColliding(r1, r2)
	collide2 = IsColliding(r2, r1)

	assert.Equal(t, true, collide1)
	assert.Equal(t, true, collide2)

	r1 = Range{Start: 1, End: 5}
	r2 = Range{Start: 1, End: 5}

	collide1 = IsColliding(r1, r2)
	collide2 = IsColliding(r2, r1)

	assert.Equal(t, true, collide1)
	assert.Equal(t, true, collide2)

	r1 = Range{Start: 1, End: 5}
	r2 = Range{Start: 6, End: 8}

	collide1 = IsColliding(r1, r2)
	collide2 = IsColliding(r2, r1)

	assert.Equal(t, false, collide1)
	assert.Equal(t, false, collide2)
}
