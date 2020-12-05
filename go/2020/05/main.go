package fifth

import (
	"fmt"
	"math"

	"github.com/tjsoler/advent-of-code/go/2020/helpers"
)

func Register() {
	fmt.Println("Day 05")
	a()
	b()
}

func a() {
	fmt.Printf("\tPart A")

	passes := helpers.ReadFileToStringSlice("./2020/05/input.txt")
	highestID := 0
	for _, pass := range passes {
		id := getBoardingPassID(pass)
		if id > highestID {
			highestID = id
		}
	}

	fmt.Printf(", solution: %d\n", highestID)
}

func b() {
	fmt.Printf("\tPart B")
	passes := helpers.ReadFileToStringSlice("./2020/05/input.txt")
	seen := make(map[int]struct{})
	for _, pass := range passes {
		seen[getBoardingPassID(pass)] = struct{}{}
	}

	highestID := 0
	for _, pass := range passes {
		id := getBoardingPassID(pass)
		if id > highestID {
			highestID = id
		}
	}
	missingID := 0
	for i := 1; i < highestID; i++ {
		_, idSeen := seen[i]
		_, hasOneMore := seen[i+1]
		_, hasOneLess := seen[i-1]
		if !idSeen && hasOneMore && hasOneLess {
			missingID = i
		}
	}
	fmt.Printf(", solution: %d\n", missingID)
}

func getBoardingPassID(boardingPass string) int {
	row := getPos(boardingPass[:7], 127, 0)
	column := getPos(boardingPass[7:], 7, 0)

	return row*8 + column
}

func getPos(boardingPass string, maxPos int, minPos int) int {
	for _, posRune := range boardingPass {
		distance := int(math.Ceil(float64(maxPos-minPos) / float64(2)))
		if rune(posRune) == 'L' || rune(posRune) == 'F' {
			maxPos = maxPos - distance
		}
		if rune(posRune) == 'R' || rune(posRune) == 'B' {
			minPos = minPos + distance
		}
	}

	return maxPos
}
