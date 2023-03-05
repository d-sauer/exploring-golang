package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"testing/fstest"
)

func TestLoad(t *testing.T) {
	tf := fstest.MapFS{"input.txt": {Data: []byte(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`)}}

	file, _ := tf.Open("input.txt")
	defer file.Close()

	count, _ := CountCollisionPairs(file)
	assert.Equal(t, 2, count)
}

func TestLoadFail(t *testing.T) {
	tf := fstest.MapFS{"input.txt": {Data: []byte(`2-4,6-
`)}}

	file, _ := tf.Open("input.txt")
	defer file.Close()

	_, err := CountCollisionPairs(file)
	assert.Error(t, err)
}
