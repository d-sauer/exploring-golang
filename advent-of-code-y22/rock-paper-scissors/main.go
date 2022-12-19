package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var fileProperty = flag.String("file", "", "Strategy file")
	flag.Parse()

	fmt.Printf("Opening file: %s\n", *fileProperty)

	file, err := os.Open(*fileProperty)
	if err != nil {
		fmt.Errorf("can't open file %s: %w", *fileProperty, err)
	}
	defer file.Close()

	sg := &StrategyGuide{}
	err = LoadStrategy(file, sg)

	if err != nil {
		fmt.Errorf("can't load strategy guide %s: %w", *fileProperty, err)
		return
	}

}
