package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readFile("./input-1.txt")
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

	fmt.Printf("Acceptable passwords: %d\n", acceptablePasswordsCount)
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
