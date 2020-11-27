package easy

import "testing"

func Test_canMakeArithmeticProgression(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case 1: + - 2",
			args: args{
				arr: []int{
					3, 5, 1,
				},
			},
			want: true,
		},
		{
			name: "test case 2: false",
			args: args{
				arr: []int{
					1, 2, 4,
				},
			},
			want: false,
		},
		{
			name: "test case 3: true",
			args: args{
				arr: []int{
					1, 100,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMakeArithmeticProgression(tt.args.arr); got != tt.want {
				t.Errorf("canMakeArithmeticProgression() = %v, want %v", got, tt.want)
			}
		})
	}
}
