package recursion

import "testing"

func Test_palindromeCheck(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test-case-1",
			args: args{
				str: "abcba",
			},
			want: true,
		},
		{
			name: "test-case-2",
			args: args{
				str: "abccba",
			},
			want: true,
		},
		{
			name: "test-case-3",
			args: args{
				str: "abcdef",
			},
			want: false,
		},
		{
			name: "test-case-4",
			args: args{
				str: "abcdefg",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindromeCheck(tt.args.str); got != tt.want {
				t.Errorf("palindromeCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
