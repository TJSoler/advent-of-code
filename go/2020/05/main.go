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

var passes = helpers.ReadFileToStringSlice("./2020/05/input.txt")

func a() {
	fmt.Printf("\tPart A")

	fmt.Printf(", solution: %d\n", getHighest(passes))
}

func b() {
	fmt.Printf("\tPart B")

	fmt.Printf(", solution: %d\n", getFirstUnused(passes))
}

func getFirstUnused(passes []string) (missing int) {
	seen := make(map[int]struct{})
	highest := getHighest(passes)

	for _, pass := range passes {
		seen[getBoardingPassID(pass)] = struct{}{}
	}

	for i := 1; i < highest+1; i++ {
		_, idSeen := seen[i]
		_, hasOneMore := seen[i+1]
		_, hasOneLess := seen[i-1]
		if !idSeen && hasOneMore && hasOneLess {
			return i
		}
	}

	return missing
}

func getHighest(passes []string) (highestID int) {
	for _, pass := range passes {
		id := getBoardingPassID(pass)
		if id > highestID {
			highestID = id
		}
	}

	return highestID
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
