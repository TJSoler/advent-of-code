package second

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tjsoler/advent-of-code/go/2020/helpers"
)

func Register() {
	fmt.Println("Day 02")
	a()
	b()
}

func a() {
	fmt.Printf("\tPart A")
	lines := helpers.ReadFileToStringSlice("./2020/02/input.txt")
	acceptablePasswordsCount := 0

	for _, line := range lines {
		// separate policy from password
		split := strings.Split(line, ":")
		password := strings.Trim(split[1], " ")

		// separate policy char from reps
		split = strings.Split(split[0], " ")
		letter := strings.Trim(split[1], " ")

		// separate min and max reps
		split = strings.Split(split[0], "-")
		min, _ := strconv.Atoi(strings.Trim(split[0], " "))
		max, _ := strconv.Atoi(strings.Trim(split[1], " "))

		appearances := strings.Count(password, letter)

		if appearances >= min && appearances <= max {
			acceptablePasswordsCount++
		}
	}

	fmt.Printf(", solution: %d\n", acceptablePasswordsCount)
}

func b() {
	fmt.Printf("\tPart B")
	lines := helpers.ReadFileToStringSlice("./2020/02/input.txt")
	acceptablePasswordsCount := 0

	for _, line := range lines {
		// separate policy from password
		split := strings.Split(line, ":")
		password := strings.Trim(split[1], " ")

		// separate policy char from reps
		split = strings.Split(split[0], " ")
		letter := rune(strings.Trim(split[1], " ")[0])

		// separate both positions
		split = strings.Split(split[0], "-")
		firstPos, _ := strconv.Atoi(strings.Trim(split[0], " "))
		secondPos, _ := strconv.Atoi(strings.Trim(split[1], " "))

		if runeAt(password, firstPos) == letter || runeAt(password, secondPos) == letter {
			acceptablePasswordsCount++
		}
		if runeAt(password, firstPos) == letter && runeAt(password, secondPos) == letter {
			acceptablePasswordsCount--
		}

	}

	fmt.Printf(", solution: %d\n", acceptablePasswordsCount)
}

func runeAt(s string, pos int) rune {
	return rune(s[pos-1])
}
