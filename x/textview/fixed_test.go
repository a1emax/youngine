package textview

import (
	"testing"

	"github.com/a1emax/youngine/basic"
)

func TestFloatToFixed(t *testing.T) {
	type args struct {
		x basic.Float
	}
	type want struct {
		result Fixed
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "0",
			args: args{
				x: 0,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "42.25",
			args: args{
				x: 42.25,
			},
			want: want{
				result: (42 << 6) + (1 << 4),
			},
		},
		{
			name: "-42.25",
			args: args{
				x: -42.25,
			},
			want: want{
				result: -((42 << 6) + (1 << 4)),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FloatToFixed(tt.args.x)
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestFixedToFloat(t *testing.T) {
	type args struct {
		x Fixed
	}
	type want struct {
		result basic.Float
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "0",
			args: args{
				x: 0,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "42.25",
			args: args{
				x: (42 << 6) + (1 << 4),
			},
			want: want{
				result: 42.25,
			},
		},
		{
			name: "-42.25",
			args: args{
				x: -((42 << 6) + (1 << 4)),
			},
			want: want{
				result: -42.25,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FixedToFloat(tt.args.x)
			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}
