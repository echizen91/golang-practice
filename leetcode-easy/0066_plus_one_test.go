package easy

/*
Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.

Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.

Input: digits = [0]
Output: [1]

** Edge Case
Input: digits = [9,9,9,9]
Output: [1,0,0,0,0]

Input: digits = [8,9,9,9]
Output: [9,0,0,0]
*/

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Input: digits = [1,2,3]",
			args: args{
				digits: []int{
					1, 2, 3,
				},
			},
			want: []int{
				1, 2, 4,
			},
		},
		{
			name: "Input: digits = [4,3,2,1]",
			args: args{
				digits: []int{
					4, 3, 2, 1,
				},
			},
			want: []int{
				4, 3, 2, 2,
			},
		},
		{
			name: "Input: digits = [0]",
			args: args{
				digits: []int{
					0,
				},
			},
			want: []int{
				1,
			},
		},
		{
			name: "Input: digits = [9,9,9,9]",
			args: args{
				digits: []int{
					9, 9, 9, 9,
				},
			},
			want: []int{
				1, 0, 0, 0, 0,
			},
		},
		{
			name: "Input: digits = [8,9,9,9]",
			args: args{
				digits: []int{
					8, 9, 9, 9,
				},
			},
			want: []int{
				9, 0, 0, 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
