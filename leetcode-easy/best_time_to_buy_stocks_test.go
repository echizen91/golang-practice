package easy

/*
Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
			 Not 7-1 = 6, as selling price needs to be larger than buying price.

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: [7,1,5,3,6,4]",
			args: args{
				prices: []int{
					7, 1, 5, 3, 6, 4,
				},
			},
			want: 5,
		},
		{
			name: "Input: [7,6,4,3,1]",
			args: args{
				prices: []int{
					7, 6, 4, 3, 1,
				},
			},
			want: 0,
		},
		{
			name: "Input: [2,1,2,1,0,1,2]",
			args: args{
				prices: []int{
					2, 1, 2, 1, 0, 1, 2,
				},
			},
			want: 2,
		},
		{
			name: "Input: [1,7,2,4]",
			args: args{
				prices: []int{
					1, 7, 2, 4,
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
