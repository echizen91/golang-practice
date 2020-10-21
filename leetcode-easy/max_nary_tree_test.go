package easy

/*
Input: root = [1,null,3,2,4,null,5,6]
Output: 3
*/

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: root = [1,null,3,2,4,null,5,6]",
			args: args{
				root: &Node{
					Val: 1,
					Children: []*Node{
						&Node{
							Val: 3,
							Children: []*Node{
								&Node{
									Val: 5,
								},
								&Node{
									Val: 6,
								},
							},
						},
						&Node{
							Val: 2,
						},
						&Node{
							Val: 4,
						},
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
