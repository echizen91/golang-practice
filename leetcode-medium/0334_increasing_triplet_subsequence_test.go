package medium

import "testing"

func Test_increasingTriplet(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test-case 1: positive true case",
			args: args{
				nums: []int{
					1, 2, 3, 4, 5,
				},
			},
			want: true,
		},
		{
			name: "test-case 2: positive mixed true case",
			args: args{
				nums: []int{
					2, 9, 4, 1, 2, 3,
				},
			},
			want: true,
		},
		{
			name: "test-case 3: negative true case",
			args: args{
				nums: []int{
					-3, -2, -1, 0, 1,
				},
			},
			want: true,
		},
		{
			name: "test-case 4: positive false case",
			args: args{
				nums: []int{
					5, 4, 3, 2, 1,
				},
			},
			want: false,
		},
		{
			name: "test-case 5: negative false case",
			args: args{
				nums: []int{
					-1, -2, -3, -4, -5,
				},
			},
			want: false,
		},
		{
			name: "test-case 6: postive same value false case",
			args: args{
				nums: []int{
					1, 1, 1, 1,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingTriplet(tt.args.nums); got != tt.want {
				t.Errorf("increasingTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}
