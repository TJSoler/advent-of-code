package third

import (
	"fmt"
	"unicode/utf8"

	"github.com/tjsoler/advent-of-code/go/2020/helpers"
)

var lines = helpers.ReadFileToStringSlice("./2020/03/input.txt")

func Register() {
	fmt.Println("Day 03")
	a()
	b()
}

func a() {
	fmt.Printf("\tPart A")
	fmt.Printf(", solution: %d\n", checkForTrees(lines, 3, 1))
}

func b() {
	fmt.Printf("\tPart B")

	slope1 := checkForTrees(lines, 1, 1)
	slope2 := checkForTrees(lines, 3, 1)
	slope3 := checkForTrees(lines, 5, 1)
	slope4 := checkForTrees(lines, 7, 1)
	slope5 := checkForTrees(lines, 1, 2)

	product := slope1 * slope2 * slope3 * slope4 * slope5

	fmt.Printf(", solution: %d\n", product)
}

func checkForTrees(lines []string, moveX int, moveY int) (treeCount int) {

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
