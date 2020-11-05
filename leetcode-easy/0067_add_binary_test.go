package easy

/*
Input: a = "11", b = "1"
Output: "100"

Input: a = "1010", b = "1011"
Output: "10101"

Input: a = "1111", b = "1111"
Output: "11110"

Input: a = "100", b = "110010"
Output: "110110"
*/

import "testing"

func TestAddBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Input: a = \"11\", b = \"1\"",
			args: args{
				a: "11",
				b: "1",
			},
			want: "100",
		},
		{
			name: "Input: a = \"1010\", b = \"1011\"",
			args: args{
				a: "1010",
				b: "1011",
			},
			want: "10101",
		},
		{
			name: "Input: a = \"1111\", b = \"1111\"",
			args: args{
				a: "1111",
				b: "1111",
			},
			want: "11110",
		},
		{
			name: "Input: a = \"100\", b = \"110010\"",
			args: args{
				a: "100",
				b: "110010",
			},
			want: "110110",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("AddBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
