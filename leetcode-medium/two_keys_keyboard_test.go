package medium

/*
Example 1:

Input: 3
Output: 3
Explanation:
Intitally, we have one character 'A'.
In step 1, we use Copy All operation.
In step 2, we use Paste operation to get 'AA'.
In step 3, we use Paste operation to get 'AAA'.

Input: 12
Output: 7

Intitally, we have one character 'A'.
1: Copy
2: Paste to get 'AA' -> 10 more
since 10%2 -> 0
3: Copy
4: Paste to get 'AAAA' -> 8 more
since 8%4 -> 0
5: Copy
6: Paste to get 'AAAAAAAA' -> 4 more
since 4<8
7: Paste to get '12As'
*/

import "testing"

func Test_minSteps(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: 3",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "Input: 12",
			args: args{
				n: 12,
			},
			want: 7,
		},
		{
			name: "Input: 20",
			args: args{
				n: 20,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSteps(tt.args.n); got != tt.want {
				t.Errorf("minSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
