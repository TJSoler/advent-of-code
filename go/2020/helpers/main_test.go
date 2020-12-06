package helpers

import (
	"reflect"
	"testing"
)

func Test_GetGroupsSeparatedByEmptyLine(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name       string
		args       args
		wantGroups []string
	}{
		{"unbroken", args{[]string{
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
			"",
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:62in"}}, []string{
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:62in"}},
		{"broken default", args{[]string{
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
			"byr:1937 iyr:2017 cid:147 hgt:183cm",
			"",
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
			"byr:1937 iyr:2017 cid:147 hgt:62in"}}, []string{
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:62in"}},
		{"super broken", args{[]string{
			"ecl:gry",
			"pid:860033327",
			"eyr:2020",
			"hcl:#fffffd",
			"byr:1937",
			"iyr:2017",
			"cid:147",
			"hgt:183cm",
			"",
			"",
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
			"byr:1937 iyr:2017 cid:147 hgt:62in"}}, []string{
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:62in"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotGroups := GetGroupsSeparatedByEmptyLine(tt.args.lines); !reflect.DeepEqual(gotGroups, tt.wantGroups) {
				t.Errorf("getGroupsSeparetedBy() = %v, want %v", gotGroups, tt.wantGroups)
			}
		})
	}
}
