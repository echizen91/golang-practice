package easy

import "testing"

func Test_maximum69Number(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-input: 6969",
			args: args{
				num: 6969,
			},
			want: 9969,
		},
		{
			name: "test-input: 6699",
			args: args{
				num: 6699,
			},
			want: 9699,
		},
		{
			name: "test-input: 9969",
			args: args{
				num: 9969,
			},
			want: 9999,
		},
		{
			name: "test-input: 9999",
			args: args{
				num: 9999,
			},
			want: 9999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum69Number(tt.args.num); got != tt.want {
				t.Errorf("maximum69Number() = %v, want %v", got, tt.want)
			}
		})
	}
}
