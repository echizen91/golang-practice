package easy

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test-case-1",
			args: args{
				s: []byte{
					'h', 'e', 'l', 'l', 'o',
				},
			},
			want: []byte{
				'o', 'l', 'l', 'e', 'h',
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
