package easy

/*
Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]

[1,0] m = 1
[2]   n = 1

[4,5,6,0,0,0]
3
[1,2,3]
3
*/

import (
	"reflect"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: `Input:
					nums1 = [1,2,3,0,0,0], m = 3
					nums2 = [2,5,6],       n = 3`,
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: `Input:
					nums1 = [1,0], m = 1
					nums2 = [2],   n = 1`,
			args: args{
				nums1: []int{1, 0},
				m:     1,
				nums2: []int{2},
				n:     1,
			},
			want: []int{1, 2},
		},
		{
			name: `Input:
					nums1 = [4,5,6,0,0,0], m = 3
					nums2 = [1,2,3],   	   n = 3`,
			args: args{
				nums1: []int{4, 5, 6, 0, 0, 0},
				m:     3,
				nums2: []int{1, 2, 3},
				n:     3,
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSortedArray(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
