package easy

/*
Input: l1 = [1,2,4], l2 = [1,3,4]
Output: [1,1,2,3,4,4]

Input: l1 = [], l2 = []
Output: []

Input: l1 = [], l2 = [0]
Output: [0]
*/

import (
	"reflect"
	"testing"
)

func TestMergeSortedList(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Input: l1 = [1,2,4], l2 = [1,3,4]",
			args: args{
				l1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
				l2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
								},
							},
						},
					},
				},
			},
		},
		// {
		// 	name: "Input: l1 = [], l2 = []",
		// 	args: args{
		// 		l1: &ListNode{},
		// 		l2: &ListNode{},
		// 	},
		// 	want: &ListNode{},
		// },
		// {
		// 	name: "Input: l1 = [], l2 = [0]",
		// 	args: args{
		// 		l1: &ListNode{},
		// 		l2: &ListNode{
		// 			Val: 0,
		// 		},
		// 	},
		// 	want: &ListNode{
		// 		Val: 0,
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSortedList(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSortedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
