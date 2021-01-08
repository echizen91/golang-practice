package easy

import "testing"

func Test_subtractProductAndSum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-case-1",
			args: args{
				n: 234,
			},
			want: 15,
		},
		{
			name: "test-case-2",
			args: args{
				n: 4421,
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtractProductAndSum(tt.args.n); got != tt.want {
				t.Errorf("subtractProductAndSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
