package fourth

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/tjsoler/advent-of-code/go/2020/helpers"
)

func Register() {
	fmt.Println("Day 04")
	a()
	b()
}

func a() {
	fmt.Printf("\tPart A")

	lines := helpers.ReadFileToStringSlice("./2020/04/input.txt")
	passportEntry := ""
	validPassports := 0

	for index, entry := range lines {
		passportEntry = passportEntry + " " + entry
		if entry == "" || index == len(lines)-1 {
			if simpleValidation(passportEntry) {
				validPassports++
			}
			passportEntry = ""
			continue
		}
	}

	fmt.Printf(", solution: %d\n", validPassports) // 222
}

func b() {
	fmt.Printf("\tPart B")
	lines := helpers.ReadFileToStringSlice("./2020/04/input.txt")
	passportEntry := ""
	validPassports := 0

	for index, entry := range lines {
		passportEntry = passportEntry + " " + entry
		if entry == "" || index == len(lines)-1 {
			if advancedValidation(passportEntry) {
				validPassports++
			}
			passportEntry = ""
			continue
		}
	}

	fmt.Printf(", solution: %d\n", validPassports)
}

func advancedValidation(passport string) bool {
	re := regexp.MustCompile(`([a-zA-Z]*):([a-zA-Z0-9#]*)`)
	matches := re.FindAllStringSubmatch(passport, -1)

	if len(matches) < 7 {
		return false
	}

	if checkYear(getFieldValue(matches, "byr"), 1920, 2002) &&
		checkYear(getFieldValue(matches, "iyr"), 2010, 2020) &&
		checkYear(getFieldValue(matches, "eyr"), 2020, 2030) &&
		checkHeight(getFieldValue(matches, "hgt")) &&
		checkHairColour(getFieldValue(matches, "hcl")) &&
		checkEyeColour(getFieldValue(matches, "ecl")) &&
		checkPassportID(getFieldValue(matches, "pid")) {
		return true
	}

	return false
}

func simpleValidation(passport string) bool {
	re := regexp.MustCompile(`([a-zA-Z]*):([a-zA-Z0-9#]*)`)
	matches := re.FindAllStringSubmatch(passport, -1)

	if len(matches) == 8 {
		return true
	}

	if len(matches) == 7 && !containsRec(matches, "cid") {
		return true
	}

	return false
}

func containsRec(s [][]string, e string) bool {
	for _, a := range s {
		if contains(a, e) {
			return true
		}
	}
	return false
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

func getFieldValue(fields [][]string, lookup string) string {
	for _, field := range fields {
		if field[1] == lookup {
			return field[2]
		}
	}

	return ""
}

func checkYear(value string, minYear int, maxYear int) bool {
	intVal, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	if intVal >= minYear && intVal <= maxYear {
		return true
	}

	return false
}

func checkHeight(value string) bool {
	var scale string
	var height int
	_, err := fmt.Sscanf(value, "%d%s", &height, &scale)
	if err != nil {
		return false
	}

	if scale == "cm" && height >= 150 && height <= 193 {
		return true
	}
	if scale == "in" && height >= 59 && height <= 76 {
		return true
	}

	return false
}

func checkHairColour(value string) bool {
	if len(value) != 7 && !strings.HasPrefix(value, "#") {
		return false
	}

	hexColor, err := hex.DecodeString(value[1:])
	if err != nil {
		return false
	}

	if between(int(hexColor[0]), 0, 255) && between(int(hexColor[1]), 0, 255) && between(int(hexColor[2]), 0, 255) {
		return true
	}
	return false
}

func checkEyeColour(value string) bool {
	possibles := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	if contains(possibles, value) {
		return true
	}
	return false
}

func between(number, min, max int) bool {
	return number >= min && number <= max
}

func checkPassportID(value string) bool {
	if len(value) != 9 {
		return false
	}

	var id int
	_, err := fmt.Sscanf(value, "%09d", &id)
	if err == nil {
		return true
	}

	return false
}
