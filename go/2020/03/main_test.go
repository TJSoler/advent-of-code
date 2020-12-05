package third

import "testing"

func Test_checkForTrees(t *testing.T) {
	type args struct {
		lines []string
		moveX int
		moveY int
	}
	tests := []struct {
		name          string
		args          args
		wantTreeCount int
	}{
		{"example", args{[]string{
			"..##.......",
			"#...#...#..",
			".#....#..#.",
			"..#.#...#.#",
			".#...##..#.",
			"..#.##.....",
			".#.#.#....#",
			".#........#",
			"#.##...#...",
			"#...##....#",
			".#..#...#.#",
		}, 3, 1}, 7},
		{"example two", args{[]string{
			"#.##.......",
			"#...#...#..",
			".#....#..#.",
			"..#.#...#.#",
			".#...##..#.",
			"..#.##.....",
			".#.#.#....#",
			".#........#",
			"#.##...#...",
			"#...##....#",
			".#..#...#.#",
		}, 3, 1}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTreeCount := checkForTrees(tt.args.lines, tt.args.moveX, tt.args.moveY); gotTreeCount != tt.wantTreeCount {
				t.Errorf("checkForTrees() = %v, want %v", gotTreeCount, tt.wantTreeCount)
			}
		})
	}
}
