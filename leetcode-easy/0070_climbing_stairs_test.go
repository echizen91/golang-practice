package easy

/*
Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/

import "testing"

func TestClimbingStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: 2",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "Input: 3",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "Input: 4",
			args: args{
				n: 4,
			},
			want: 5,
		},
		{
			name: "Input: 6",
			args: args{
				n: 6,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClimbingStairs(tt.args.n); got != tt.want {
				t.Errorf("ClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
