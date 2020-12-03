package third

import (
	"fmt"
	"unicode/utf8"

	"github.com/tjsoler/advent-of-code/go/2020/helpers"
)

func Register() {
	fmt.Println("Day 03")
	a()
	b()
}

func a() {
	fmt.Printf("\tPart A")
	fmt.Printf(", solution: %d\n", checkForTrees(3, 1))
}

func b() {
	fmt.Printf("\tPart B")

	slope1 := checkForTrees(1, 1)
	slope2 := checkForTrees(3, 1)
	slope3 := checkForTrees(5, 1)
	slope4 := checkForTrees(7, 1)
	slope5 := checkForTrees(1, 2)

	product := slope1 * slope2 * slope3 * slope4 * slope5

	fmt.Printf(", solution: %d\n", product)
}

func checkForTrees(moveX int, moveY int) (treeCount int) {
	lines := helpers.ReadFileToStringSlice("./2020/03/input.txt")
	x, y := 0, 0
	tree := '#'

	for y <= len(lines)-1 {
		step := lines[y]
		tempX := x
		for tempX > utf8.RuneCountInString(step)-1 {
			tempX = tempX - utf8.RuneCountInString(step)
		}

		if rune(step[tempX]) == tree {
			treeCount++
		}

		x = x + moveX
		y = y + moveY
	}

	return treeCount
}
