package easy

/*
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Input: nums = [1]
Output: 1

Input: nums = [0]
Output: 0

Input: nums = [-1]
Output: -1

Input: nums = [-2147483647]
Output: -2147483647
*/

import "testing"

func TestMaxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: nums = [-2,1,-3,4,-1,2,1,-5,4]",
			args: args{
				nums: []int{
					-2, 1, -3, 4, -1, 2, 1, -5, 4,
				},
			},
			want: 6,
		},
		{
			name: "Input: nums = [1]",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "Input: nums = [0]",
			args: args{
				nums: []int{0},
			},
			want: 0,
		},
		{
			name: "Input: nums = [-1]",
			args: args{
				nums: []int{-1},
			},
			want: -1,
		},
		{
			name: "Input: nums = [-2147483647]",
			args: args{
				nums: []int{-2147483647},
			},
			want: -2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("MaxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
