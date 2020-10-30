package geekmath

import "testing"

func Test_isPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "(base case) test input = 1",
			args: args{
				n: 1,
			},
			want: false,
		},
		{
			name: "(base case) test input = 2",
			args: args{
				n: 2,
			},
			want: true,
		},
		{
			name: "(base case) test input = 3",
			args: args{
				n: 3,
			},
			want: true,
		},
		{
			name: "test input = 4",
			args: args{
				n: 4,
			},
			want: false,
		},
		{
			name: "test input = 17",
			args: args{
				n: 17,
			},
			want: true,
		},
		{
			name: "test input = 61",
			args: args{
				n: 61,
			},
			want: true,
		},
		{
			name: "test input = 100",
			args: args{
				n: 100,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.args.n); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
