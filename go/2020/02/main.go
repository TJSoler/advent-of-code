package second

import (
	"fmt"
	"strings"

	"github.com/tjsoler/advent-of-code/go/2020/helpers"
)

var lines = helpers.ReadFileToStringSlice("./2020/02/input.txt")

func Register() {
	fmt.Println("Day 02")
	a()
	b()
}

func a() {
	fmt.Printf("\tPart A")

	acceptablePasswordsCount := getAcceptablePasswordsCount(lines, countValidator)

	fmt.Printf(", solution: %d\n", acceptablePasswordsCount)
}

func b() {
	fmt.Printf("\tPart B")

	acceptablePasswordsCount := getAcceptablePasswordsCount(lines, posValidator)

	fmt.Printf(", solution: %d\n", acceptablePasswordsCount)
}

type validator func(int, int, string, string) bool

func countValidator(min int, max int, letter string, password string) bool {
	appearances := strings.Count(password, letter)

	if appearances >= min && appearances <= max {
		return true
	}
	return false
}

func posValidator(firstPos int, secondPos int, letter string, password string) bool {
	if secondPos > len(password) {
		return false
	}

	if runeAt(password, firstPos) == rune(letter[0]) && runeAt(password, secondPos) == rune(letter[0]) {
		return false
	}

	if runeAt(password, firstPos) == rune(letter[0]) || runeAt(password, secondPos) == rune(letter[0]) {
		return true
	}

	return false
}

func getAcceptablePasswordsCount(lines []string, validatorfn validator) (acceptablePasswordsCount int) {
	for _, line := range lines {
		var password string
		var letter string
		var min, max int

		_, err := fmt.Sscanf(line, "%d-%d %1s: %s", &min, &max, &letter, &password)
		if err != nil {
			fmt.Println("error parsing passwords with scanf")
			return 0
		}

		if validatorfn(min, max, letter, password) {
			acceptablePasswordsCount++
		}
	}

	return acceptablePasswordsCount
}

func runeAt(s string, pos int) rune {
	return rune(s[pos-1])
}
