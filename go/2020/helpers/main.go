package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFileToStringSlice(filePath string) (lines []string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ReadFileToIntSlice(filePath string) (lines []int) {
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

func GetGroupsSeparatedByEmptyLine(lines []string) (groups []string) {
	groupPiece := ""

	for index, piece := range lines {
		groupPiece = groupPiece + " " + piece
		if index == len(lines)-1 {
			piece = ""
		}

		groupPiece = strings.TrimSpace(groupPiece)
		if len(piece) == 0 && len(groupPiece) > 0 {
			groups = append(groups, groupPiece)
			groupPiece = ""
			continue
		}
	}

	return groups
}
