package easy

/*
Input: 4
Output: 2

Input: 8
Output: 2

Input: 10
Output: 3
*/

import "testing"

func TestSquareRoot(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: 4",
			args: args{
				x: 4,
			},
			want: 2,
		},
		{
			name: "Input: 8",
			args: args{
				x: 8,
			},
			want: 2,
		},
		{
			name: "Input: 10",
			args: args{
				x: 10,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SquareRoot(tt.args.x); got != tt.want {
				t.Errorf("SquareRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
