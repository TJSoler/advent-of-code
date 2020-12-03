package first

import (
	"fmt"

	"github.com/tjsoler/advent-of-code/go/2020/helpers"
)

func Register() {
	fmt.Println("Day 01")
	a()
	b()
}

func a() {
	fmt.Printf("\tPart A")
	target := 2020
	lines := helpers.ReadFileToIntSlice("./2020/01/input.txt")
	seen := make(map[int]struct{})

	for _, number := range lines {
		seen[number] = struct{}{}
		_, ok := seen[target-number]
		if !ok {
			fmt.Printf(", solution: %d\n", number*(target-number))

			break
		}
	}
}

func b() {
	fmt.Printf("\tPart B")
	lines := helpers.ReadFileToIntSlice("./2020/01/input.txt")
	for _, line := range lines {
		for _, secondLine := range lines {
			for _, thirdLine := range lines {
				if line+secondLine+thirdLine == 2020 {
					fmt.Printf(", solution: %d\n", line*secondLine*thirdLine)
					return
				}
			}
		}

	}
}
