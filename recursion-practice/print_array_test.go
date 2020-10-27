package recursion

import "testing"

func Test_printArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-case-1",
			args: args{
				nums: []int{
					2, 4, 6, 8, 1,
				},
			},
			want: "2 4 6 8 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printArray(tt.args.nums); got != tt.want {
				t.Errorf("printArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
