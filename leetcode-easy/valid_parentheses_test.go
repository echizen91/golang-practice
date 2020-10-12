package easy

/*
Input: s = "()"
Output: true

Input: s = "()[]{}"
Output: true

Input: s = "(]"
Output: false

Input: s = "([)]"
Output: false

Input: s = "{[]}"
Output: true

Constraints:
1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

import "testing"

func TestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Input: s = ()",
			args: args{s: "()"},
			want: true,
		},
		{
			name: "Input: s = ()[]{}",
			args: args{s: "()[]{}"},
			want: true,
		},
		{
			name: "Input: s = (]",
			args: args{s: "(]"},
			want: false,
		},
		{
			name: "Input: s = ([)]",
			args: args{s: "([)]"},
			want: false,
		},
		{
			name: "Input: s = {[]}",
			args: args{s: "{[]}"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("ValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
