package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

const (
	MY_ROCK          = 'X'
	MY_PAPER         = 'Y'
	MY_SCISSOR       = 'Z'
	OPPONENT_ROCK    = 'A'
	OPPONENT_PAPER   = 'B'
	OPPONENT_SCISSOR = 'C'
)

func main() {
	file, err := os.Open("day02.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	lines := 0
	for scanner.Scan() {
		line := scanner.Text()
		lines++

		// 1 for ROCK -> X
		// 2 for PAPER -> Y
		// 3 for SCISSORS -> Z

		// R P S -> A B C

		// calculation for part one
		point := 0
		opponent := line[0]
		mine := line[2]

		if mine == MY_ROCK {
			point += 1

			if opponent == OPPONENT_SCISSOR {
				point += 6
			}
		} else if mine == MY_PAPER {
			point += 2

			if opponent == OPPONENT_ROCK {
				point += 6
			}
		} else if mine == MY_SCISSOR {
			point += 3

			if opponent == OPPONENT_PAPER {
				point += 6
			}
		}

		if (mine - 'X') == (opponent - 'A') {
			total += 3
		}

		total += point

		// calculation for second part

		// mine:
		// X -> Lose
		// Y -> Draw
		// Z -> Win

		point = 0
		if mine == 'X' {

		} else if mine == 'Y' {

		} else if mine == 'Z' {

		}

	}

	fmt.Println("Part one: " + fmt.Sprint(total))
}
