package easy

/*
Given nums = [3,2,2,3], val = 3,
Your function should return length = 2, with the first two elements of nums being 2.
It doesn't matter what you leave beyond the returned length.

Given nums = [0,1,2,2,3,0,4,2], val = 2,
Your function should return length = 5, with the first five elements of nums containing 0, 1, 3, 0, and 4.
Note that the order of those five elements can be arbitrary.

It doesn't matter what values are set beyond the returned length.

Solution:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Element.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Remove Element.
*/

import "testing"

func TestRemoveElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Given nums = [3,2,2,3]",
			args: args{
				val: 3,
				nums: []int{
					3, 2, 2, 3,
				},
			},
			want: 2,
		},
		{
			name: "Given nums = [0,1,2,2,3,0,4,2]",
			args: args{
				val: 2,
				nums: []int{
					0, 1, 2, 2, 3, 0, 4, 2,
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
