package recursion

import "testing"

func Test_recursivePow(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-case-1",
			args: args{
				a: 2,
				b: 5,
			},
			want: 32,
		},
		{
			name: "test-case-2",
			args: args{
				a: 3,
				b: 4,
			},
			want: 81,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recursivePow(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("recursivePow() = %v, want %v", got, tt.want)
			}
		})
	}
}
