package main

import (
	"fmt"
)

func check(player int, index int, arena []int) bool {
	result := false
	switch index {
	case 0:
		if arena[1] == player && arena[2] == player {
			result = true
			break
		}
		
		if arena[3] == player && arena[6] == player {
			result = true
			break
		}

		if arena[4] == player && arena[8] == player {
			result = true
			break
		}
	}

	return result
}

func main() {
	turn := true
	sum := 0
	winner := 0

	arena := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for ; sum < 9; {
		sum++
	}
 }