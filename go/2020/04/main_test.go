package fourth

import (
	"testing"
)

func Test_simpleValidation(t *testing.T) {
	tests := []struct {
		name     string
		passport string
		want     bool
	}{
		{"valid", "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm", true},
		{"invalid", "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929", false},
		{"valid-cheat", "hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm", true},
		{"extrainvalid", "hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simpleValidation(tt.passport); got != tt.want {
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkYear(t *testing.T) {
	type args struct {
		value   string
		minYear int
		maxYear int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid birth year", args{"2002", 1920, 2002}, true},
		{"invalid birth year", args{"2003", 1920, 2002}, false},
		{"errored birth year", args{"as23", 1920, 2002}, false},
		{"invalid birth year", args{"1919", 1920, 2002}, false},
		{"valid issue year", args{"2011", 2010, 2020}, true},
		{"invalid issue year", args{"2009", 2010, 2020}, false},
		{"errored issue year", args{"fffsd", 2010, 2020}, false},
		{"invalid issue year", args{"2021", 2010, 2020}, false},
		{"valid expiration year", args{"2021", 2020, 2030}, true},
		{"invalid expiration year", args{"2019", 2020, 2030}, false},
		{"errored expiration year", args{"afg11!3", 2020, 2030}, false},
		{"invalid expiration year", args{"2031", 2020, 2030}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkYear(tt.args.value, tt.args.minYear, tt.args.maxYear); got != tt.want {
				t.Errorf("checkYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_advancedValidation(t *testing.T) {
	tests := []struct {
		name     string
		passport string
		want     bool
	}{
		{"valid", "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm", true},
		{"valid inches", "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:62in", true},
		{"invalid scale", "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:58sc", false},
		{"invalid", "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929", false},
		{"valid cheat", "hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm", true},
		{"extra invalid", "hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := advancedValidation(tt.passport); got != tt.want {
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPassportID(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid", args{"012345678"}, true},
		{"invalid length", args{"12345678"}, false},
		{"invalid chars", args{"a12345678"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPassportID(tt.args.value); got != tt.want {
				t.Errorf("checkPassportID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkEyeColour(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid", args{"amb"}, true},
		{"invalid", args{"nop"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkEyeColour(tt.args.value); got != tt.want {
				t.Errorf("checkEyeColour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkHairColour(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid", args{"#AA12CC"}, true},
		{"valid only numbers", args{"#123456"}, true},
		{"valid only letters", args{"#AABBCC"}, true},
		{"invalid no pound", args{"AABBCC"}, false},
		{"invalid no hex", args{"#ZZZZZZ"}, false},
		{"invalid short", args{"#123"}, false},
		{"invalid long", args{"#123456789"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkHairColour(tt.args.value); got != tt.want {
				t.Errorf("checkHairColour() = %v, want %v", got, tt.want)
			}
		})
	}
}
