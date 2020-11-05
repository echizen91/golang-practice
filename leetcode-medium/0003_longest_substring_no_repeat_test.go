package medium

/*
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Input: s = ""
Output: 0
*/

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: s = \"abcabcbb\"",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "Input: s = \"bbbbb\"",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "Input: s = \"pwwkew\"",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "Input: s = \"\"",
			args: args{
				s: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
