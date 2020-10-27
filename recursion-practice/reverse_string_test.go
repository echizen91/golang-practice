package recursion

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-case-1",
			args: args{
				str: "abcdef",
			},
			want: "fedcba",
		},
		{
			name: "test-case-2",
			args: args{
				str: "abcdefg",
			},
			want: "gfedcba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseString(tt.args.str); got != tt.want {
				t.Errorf("reverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
