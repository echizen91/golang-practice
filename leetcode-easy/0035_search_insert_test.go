package easy

/*
Input: nums = [1,3,5,6], target = 5
Output: 2

Input: nums = [1,3,5,6], target = 2
Output: 1

Input: nums = [1,3,5,6], target = 7
Output: 4

Input: nums = [1,3,5,6], target = 0
Output: 0

Input: nums = [1], target = 0
Output: 0
*/

import "testing"

func TestSearchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: nums = [1,3,5,6], target = 5",
			args: args{
				nums: []int{
					1, 3, 5, 6,
				},
				target: 5,
			},
			want: 2,
		},
		{
			name: "Input: nums = [1,3,5,6], target = 2",
			args: args{
				nums: []int{
					1, 3, 5, 6,
				},
				target: 2,
			},
			want: 1,
		},
		{
			name: "Input: nums = [1,3,5,6], target = 7",
			args: args{
				nums: []int{
					1, 3, 5, 6,
				},
				target: 7,
			},
			want: 4,
		},
		{
			name: "Input: nums = [1,3,5,6], target = 0",
			args: args{
				nums: []int{
					1, 3, 5, 6,
				},
				target: 0,
			},
			want: 0,
		},
		{
			name: "Input: nums = [1], target = 0",
			args: args{
				nums: []int{
					1,
				},
				target: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SearchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
