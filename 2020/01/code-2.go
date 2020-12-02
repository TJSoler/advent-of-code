package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines := readFile("./input-1.txt")
	for index, line := range lines {
		for _, secondLine := range lines {
			for _, thirdLine := range lines {
				if line+secondLine+thirdLine != 2020 {
					continue
				}

				fmt.Println("result found:")
				fmt.Println(line * secondLine * thirdLine)
			}
		}
		lines = append(lines[:index], lines[index+1:]...)
	}
}

func readFile(filePath string) (lines []int) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineInt, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			return nil
		}
		lines = append(lines, lineInt)
	}
	return lines
}
