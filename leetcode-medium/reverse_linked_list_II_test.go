package medium

import (
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head *ListNode
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test-case: m=2, n=4",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
				m: 2,
				n: 4,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
		{
			name: "test-case: m=1,2 with 3 ll",
			args: args{
				head: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 2,
						},
					},
				},
				m: 2,
				n: 3,
			},
			want: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
		{
			name: "test-case: m=1, n=2",
			args: args{
				head: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
					},
				},
				m: 1,
				n: 2,
			},
			want: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
		{
			name: "test-case: m=1, n=3",
			args: args{
				head: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 7,
						},
					},
				},
				m: 1,
				n: 3,
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
