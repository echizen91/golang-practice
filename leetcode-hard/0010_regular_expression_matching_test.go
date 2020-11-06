package hard

/*
Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".

Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".

Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".

Input: s = "aab", p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".

Input: s = "mississippi", p = "mis*is*p*."
Output: false
*/

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "False output, length of p < s",
			args: args{
				s: "aaa",
				p: "a*a",
			},
			want: true,
		},
		{
			name: "True output, p with *",
			args: args{
				s: "aa",
				p: "a*",
			},
			want: true,
		},
		{
			name: "True output, wildcard + *",
			args: args{
				s: "ab",
				p: ".*",
			},
			want: true,
		},
		{
			name: "True output, * checks",
			args: args{
				s: "aab",
				p: "c*a*b",
			},
			want: true,
		},
		{
			name: "False output, long string with *",
			args: args{
				s: "mississippi",
				p: "mis*is*p*.",
			},
			want: false,
		},
		{
			name: "False output, extra char in p",
			args: args{
				s: "aab",
				p: "ca*b",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
