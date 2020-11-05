package easy

/*
Input: 1->1->2
Output: 1->2

Input: 1->1->2->3->3
Output: 1->2->3
*/

import (
	"reflect"
	"testing"
)

func TestRemoveDupSortedList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Input: Empty",
			args: args{
				head: &ListNode{},
			},
			want: &ListNode{},
		},
		{
			name: "Input: 1->1->2",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
		{
			name: "Input: 1->1->2->3->3",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 3,
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDupSortedList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDupSortedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
