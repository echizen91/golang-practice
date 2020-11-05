package easy

/*
Given nums = [1,1,2],
Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.
It doesn't matter what you leave beyond the returned length.

Given nums = [0,0,1,1,1,2,2,3,3,4],
Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.
It doesn't matter what values are set beyond the returned length.
*/

import "testing"

func TestArrayRemoveDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Given nums = [0,0,1,1,1,2,2,3,3,4]",
			args: args{
				nums: []int{
					0, 0, 1, 1, 1, 2, 2, 3, 3, 4,
				},
			},
			want: 5,
		},
		{
			name: "Given nums = [1,1,2]",
			args: args{
				nums: []int{
					1, 1, 2,
				},
			},
			want: 2,
		},
		{
			name: "Given nums = [1]",
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
			if got := ArrayRemoveDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("ArrayRemoveDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
