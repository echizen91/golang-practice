package easy

/*
Input: nums = [2,2,1]
Output: 1

Input: nums = [4,1,2,1,2]
Output: 4

Input: nums = [1]
Output: 1
*/

import "testing"

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: nums = [2,2,1]",
			args: args{
				nums: []int{
					2, 2, 1,
				},
			},
			want: 1,
		},
		{
			name: "Input: nums = [4,1,2,1,2]",
			args: args{
				nums: []int{
					4, 1, 2, 1, 2,
				},
			},
			want: 4,
		},
		{
			name: "Input: nums = [1]",
			args: args{
				nums: []int{
					1,
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
