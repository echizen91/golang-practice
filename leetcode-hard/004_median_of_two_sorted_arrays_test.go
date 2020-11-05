package hard

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test-case: odd length, different numbers",
			args: args{
				nums1: []int{
					1, 3,
				},
				nums2: []int{
					2,
				},
			},
			want: 2.00000,
		},
		{
			name: "test-case: even length, different numbers",
			args: args{
				nums1: []int{
					1, 3, 4,
				},
				nums2: []int{
					2,
				},
			},
			want: 2.50000,
		},
		{
			name: "test-case: even length, all same numbers",
			args: args{
				nums1: []int{
					0, 0,
				},
				nums2: []int{
					0, 0,
				},
			},
			want: 0.00000,
		},
		{
			name: "test-case: even length, 2 only",
			args: args{
				nums1: []int{
					100001,
				},
				nums2: []int{
					100000,
				},
			},
			want: 100000.50000,
		},
		{
			name: "test-case: odd length, 1 only",
			args: args{
				nums1: []int{
					100001,
				},
			},
			want: 100001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
