package sixth

import (
	"testing"
)

func Test_getBoardingPassID(t *testing.T) {
	type args struct {
		boardingPass string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{"FBFBBFFRLR"}, 357},
		{"example 2", args{"BFFFBBFRRR"}, 567},
		{"example 3", args{"FFFBBBFRRR"}, 119},
		{"example 4", args{"BBFFBBFRLL"}, 820},
		{"example 5", args{"FFFFFFFLLL"}, 0},
		{"example 6", args{"FFFFFFFLLR"}, 1},
		{"example 6", args{"FFFFFFFLRR"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBoardingPassID(tt.args.boardingPass); got != tt.want {
				t.Errorf("getBoardingPassID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPos(t *testing.T) {
	type args struct {
		boardingPass string
		maxPos       int
		minPos       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Row Example 1", args{"FBFBBFF", 127, 0}, 44},
		{"Column Example 1", args{"RLR", 7, 0}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPos(tt.args.boardingPass, tt.args.maxPos, tt.args.minPos); got != tt.want {
				t.Errorf("getPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHighest(t *testing.T) {
	type args struct {
		passes []string
	}
	tests := []struct {
		name          string
		args          args
		wantHighestID int
	}{
		{"test", args{[]string{"BBFBBBBRRL", "FBFFFFBLRL", "FBFBBFFRLR"}}, 894},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHighestID := getHighest(tt.args.passes); gotHighestID != tt.wantHighestID {
				t.Errorf("getHighest() = %v, want %v", gotHighestID, tt.wantHighestID)
			}
		})
	}
}

func Test_getFirstUnused(t *testing.T) {
	type args struct {
		passes []string
	}
	tests := []struct {
		name        string
		args        args
		wantMissing int
	}{
		{"test", args{[]string{"BBFBBBBRRL", "FBFFFFBLRL", "FBFBBFFRLR"}}, 0},
		{"test 2", args{[]string{"FFFFFFFLLL", "FFFFFFFLLR", "FFFFFFFLRR"}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMissing := getFirstUnused(tt.args.passes); gotMissing != tt.wantMissing {
				t.Errorf("getFirstUnused() = %v, want %v", gotMissing, tt.wantMissing)
			}
		})
	}
}
