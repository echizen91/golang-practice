package hard

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "not rotated",
			args: args{
				nums: []int{
					1, 3, 5,
				},
			},
			want: 1,
		},
		{
			name: "3 duplicates",
			args: args{
				nums: []int{
					2, 2, 2, 0, 1,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
