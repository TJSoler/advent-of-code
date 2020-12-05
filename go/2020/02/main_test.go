package second

import (
	"testing"
)

func Test_runeAt(t *testing.T) {
	type args struct {
		s   string
		pos int
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"valid", args{"example", 3}, 'a'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runeAt(tt.args.s, tt.args.pos); got != tt.want {
				t.Errorf("runeAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAcceptablePasswordsCount(t *testing.T) {
	type args struct {
		lines       []string
		validatorfn validator
	}
	tests := []struct {
		name                         string
		args                         args
		wantAcceptablePasswordsCount int
	}{
		{"all valid count", args{[]string{"1-3 a: abcde", "1-3 b: cdbefg", "2-9 c: ccccccccc"}, countValidator}, 3},
		{"all invalid count", args{[]string{"1-3 a: bcde", "1-3 b: cdefg", "2-9 z: ccccccccc"}, countValidator}, 0},
		{"mixed count", args{[]string{"1-3 a: abcde", "1-3 b: cdbefg", "2-9 z: ccccccccc"}, countValidator}, 2},
		{"all valid position", args{[]string{"1-3 a: abcde", "1-3 b: cdbefg", "2-9 c: cccccccaa"}, posValidator}, 3},
		{"all invalid position", args{[]string{"1-3 a: bcde", "1-3 b: cdefg", "2-9 z: ccccccccc"}, posValidator}, 0},
		{"mixed position", args{[]string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 z: ccccccccc"}, posValidator}, 1},
		{"lenght check position", args{[]string{"2-9 z: cc"}, posValidator}, 0},
		{"both right makes it wrong position", args{[]string{"2-9 c: ccccccccc"}, posValidator}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAcceptablePasswordsCount := getAcceptablePasswordsCount(tt.args.lines, tt.args.validatorfn); gotAcceptablePasswordsCount != tt.wantAcceptablePasswordsCount {
				t.Errorf("getAcceptablePasswordsCount() = %v, want %v", gotAcceptablePasswordsCount, tt.wantAcceptablePasswordsCount)
			}
		})
	}
}
