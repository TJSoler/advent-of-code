package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readFile("./input-2.txt")
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
			fmt.Printf("XX - ")
			acceptablePasswordsCount++
		}
		if runeAt(password, firstPos) == letter && runeAt(password, secondPos) == letter {
			acceptablePasswordsCount--
		}

		fmt.Printf("[%s] - %s - %03d-%03d\n", password, string(letter), firstPos, secondPos)
	}

	fmt.Printf("Acceptable passwords: %d\n", acceptablePasswordsCount)
}

func runeAt(s string, pos int) rune {
	return rune(s[pos-1])
}

func readFile(filePath string) (lines []string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		lines = append(lines, scanner.Text())
	}
	return lines
}
