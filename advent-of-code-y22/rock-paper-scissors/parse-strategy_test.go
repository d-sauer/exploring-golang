package main

import (
	"testing"
	"testing/fstest"
)

func TestLoad(t *testing.T) {
	//	tf := fstest.MapFS{"strategy.txt": {Data: []byte(`A Y
	//B X
	//C Z
	//`)}}

	tf := fstest.MapFS{"strategy.txt": {Data: []byte(`A X
A Y
A Z
B X
B Y
B Z
C X
C Y
C Z
`)}}

	file, err := tf.Open("strategy.txt")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	sg := &StrategyGuide{}
	err = LoadStrategy(file, sg)

}
