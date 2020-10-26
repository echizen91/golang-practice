package medium

/*
Input:  [1,2,3,4]
Output: [24,12,8,6]

Input:  [3,6,2,5]
Output: [60,30,90,36]

*/

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test-case-1",
			args: args{
				nums: []int{
					1, 2, 3, 4,
				},
			},
			want: []int{
				24, 12, 8, 6,
			},
		},
		{
			name: "test-case-2",
			args: args{
				nums: []int{
					3, 6, 2, 5,
				},
			},
			want: []int{
				60, 30, 90, 36,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
