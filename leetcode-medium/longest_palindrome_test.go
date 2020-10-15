package medium

/*
Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Input: s = "cbbd"
Output: "bb"

Input: s = "a"
Output: "a"

Input: s = "ac"
Output: "a"

Input: s = "eabcb"
Output: "bcb"
*/

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Input: s = \"babad\"",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "Input: s = \"cbbd\"",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "Input: s = \"a\"",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "Input: s = \"ac\"",
			args: args{
				s: "ac",
			},
			want: "a",
		},
		{
			name: "Input: s = \"eabcb\"",
			args: args{
				s: "eabcb",
			},
			want: "bcb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
