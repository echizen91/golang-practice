package medium

/*
Input: m = 3, n = 7
Output: 28

Input: m = 3, n = 2
Output: 3

Input: m = 7, n = 3
Output: 28

Input: m = 3, n = 3
Output: 6
*/

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: m = 3, n = 7",
			args: args{
				m: 3,
				n: 7,
			},
			want: 28,
		},
		{
			name: "Input: m = 3, n = 2",
			args: args{
				m: 3,
				n: 2,
			},
			want: 3,
		},
		{
			name: "Input: m = 3, n = 7",
			args: args{
				m: 7,
				n: 3,
			},
			want: 28,
		},
		{
			name: "Input: m = 3, n = 7",
			args: args{
				m: 3,
				n: 3,
			},
			want: 6,
		},
		{
			name: "Input: m = 23, n = 12",
			args: args{
				m: 23,
				n: 12,
			},
			want: 193536720,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
