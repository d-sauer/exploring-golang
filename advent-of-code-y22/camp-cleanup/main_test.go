package main

import (
	"fmt"
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

	file, err := tf.Open("input.txt")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	count, _ := CountCollisionPairs(file)
	fmt.Println(count)
}
