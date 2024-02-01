package main

import (
	"fmt"
	"math/rand"
	"time"
)

func check(player int, index int, arena []int) bool {
	result := false
	winPatterns := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Rows
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Columns
		{0, 4, 8}, {2, 4, 6}, // Diagonal
	}

	for _, pattern := range winPatterns {
		flag := false
		for _, idx := range pattern {
			if idx == index {
				flag = true
				break
			}
		}

		if flag {
			for _, occupation := range pattern {
				if arena[occupation] != player {
					flag = false
					break
				}
			}

			if flag {
				result = true
			}
		}

		if result {
			break
		}
	}

	return result
}

func printArena(arena []int) {
	fmt.Println(" +---+---+---+")
	for idx, cell := range arena {
		var c byte
		if cell == 1 {
			c = 'X'
		} else if cell == 2 {
			c = 'O'
		} else {
			c = ' '
		}

		fmt.Printf(" | %c",c)
		if idx % 3 == 2{
			fmt.Print(" |\n")
			fmt.Println(" +---+---+---+")
		}
	}
}

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
		winner := 0

		for j := 1; j <= turns; j++ {
			// Simulate the player's move by choosing a random, unoccupied index on the tic-tac-toe board.
			// Seed the random number generator with the current UnixNano time to ensure randomness.
			// Loop until a valid, unoccupied index is found, breaking out of the loop when such an index is obtained.
			rand.Seed(time.Now().UnixNano())
			randomIndex := rand.Intn(9)
			for {
				randomIndex = rand.Intn(9)
				if arena[randomIndex] == 0 {
					break
				}
			}

			// Player makes a move by occupying an unoccupied cell in the arena.
			arena[randomIndex] = player

			if j >= 5 {
				isWin := check(player, randomIndex, arena)
				if isWin {
					winner = player
					break
				}
			}

			if player % 2 == 1 {
				player = 2
			} else {
				player = 1
			}
		}

		if winner != 0 {
			fmt.Println(" Player", player, "Is the winner")
		} else {
			fmt.Println(" DRAW")
		}
		printArena(arena)
		fmt.Println("")
	}

}