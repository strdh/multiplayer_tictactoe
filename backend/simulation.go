package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func check(player int, index int, arena []int) bool {
// 	result := false
// 	return result
// }

func main() {
	rounds := 10

	var player int
	for i := 1; i <= rounds; i++ {
		if i % 2 == 1 {
			player = 1
		} else {
			player = 2
		}

		arena := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
		turns := 9

		for j := 1; j <= turns; j++ {
			rand.Seed(time.Now().UnixNano())
			randomIndex := rand.Intn(9)

			for {
				randomIndex = rand.Intn(9)
				if arena[randomIndex] == 0 {
					break
				}
			}

			arena[randomIndex] = player

			if player % 2 == 1 {
				player = 2
			} else {
				player = 1
			}
		}

		fmt.Println(arena)
	}
}