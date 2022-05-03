package sixth

import (
	"fmt"
	"unicode"

	"github.com/tjsoler/advent-of-code/go/2020/helpers"
)

func Register() {
	fmt.Println("Day 06")
	a()
	b()
}

var answers = helpers.ReadFileToStringSlice("./2020/06/input.txt")

func a() {
	fmt.Printf("\tPart A")

	fmt.Printf(", solution: %d\n", getSumOfAffirmativeAnswers(answers)) // 7257 is too high
}

func b() {
	fmt.Printf("\tPart B")

	fmt.Printf(", solution: %d\n", 0)
}

func getSumOfAffirmativeAnswers(lines []string) (count int) {
	answers := helpers.GetGroupsSeparatedByEmptyLine(lines)
	for _, answerGroup := range answers {
		seen := make(map[rune]struct{})
		for _, answer := range answerGroup {
			if unicode.IsSpace(answer) {
				continue
			}

			seen[answer] = struct{}{}
		}
		count = count + len(seen)
	}

	return count
}
