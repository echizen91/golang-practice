package medium

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
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
					2, 0, 2, 1, 1, 0,
				},
			},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "test-case-2",
			args: args{
				nums: []int{
					2, 0, 1,
				},
			},
			want: []int{0, 1, 2},
		},
		{
			name: "test-case-3",
			args: args{
				nums: []int{
					0,
				},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortColors(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortColors() = %v, want %v", got, tt.want)
			}
		})
	}
}
