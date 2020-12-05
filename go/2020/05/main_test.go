package fifth

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
