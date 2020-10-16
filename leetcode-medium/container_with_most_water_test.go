package medium

/*
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Input: height = [1,1]
Output: 1

Input: height = [4,3,2,1,4]
Output: 16

Input: height = [1,2,1]
Output: 2
*/

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: height = [1,8,6,2,5,4,8,3,7]",
			args: args{
				height: []int{
					1, 8, 6, 2, 5, 4, 8, 3, 7,
				},
			},
			want: 49,
		},
		{
			name: "Input: height = [1,1]",
			args: args{
				height: []int{
					1, 1,
				},
			},
			want: 1,
		},
		{
			name: "Input: height = [4,3,2,1,4]",
			args: args{
				height: []int{
					4, 3, 2, 1, 4,
				},
			},
			want: 16,
		},
		{
			name: "Input: height = [1,2,1]",
			args: args{
				height: []int{
					1, 2, 1,
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
