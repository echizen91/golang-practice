package easy

/*
Input: [3,9,20,null,null,15,7]
Output:
[
  [15,7],
  [9,20],
  [3]
]
*/

import (
	"reflect"
	"testing"
)

func TestLevelOrderBottom(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// {
		// 	name: "Input: [3,9,20,null,null,15,7]",
		// 	args: args{
		// 		root: &TreeNode{
		// 			Val: 3,
		// 			Left: &TreeNode{
		// 				Val: 9,
		// 			},
		// 			Right: &TreeNode{
		// 				Val: 20,
		// 				Left: &TreeNode{
		// 					Val: 15,
		// 				},
		// 				Right: &TreeNode{
		// 					Val: 7,
		// 				},
		// 			},
		// 		},
		// 	},
		// 	want: [][]int{
		// 		{15, 17},
		// 		{9, 20},
		// 		{3},
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LevelOrderBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelOrderBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
