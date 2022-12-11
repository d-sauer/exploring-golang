package calories

import (
	"bufio"
	"fmt"
	"io/fs"
	"strconv"
	"strings"
)

type Inventory interface {
	Process(groupId int, groupSum int)
}

func LoadInventory(file fs.File, inventory *InventoryIndex) error {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var line = 0
	var groupSum = 0
	var groupCount = 1
	var startNewGroup = true
	for fileScanner.Scan() {
		line++
		content := strings.TrimSpace(fileScanner.Text())

		if len(content) == 0 {
			inventory.Process(groupCount, groupSum)
			groupSum = 0
			groupCount++
			startNewGroup = true
			continue
		}

		startNewGroup = false
		calories, err := strconv.Atoi(content)
		if err != nil {
			return fmt.Errorf("can't parse number at line %d, due to error: %w", line, err)
		}
		groupSum += calories
	}

	if startNewGroup == false {
		inventory.Process(groupCount, groupSum)
	}

	return nil
}
