package safeuint

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		x uint
		y uint
	}
	type want struct {
		result   uint
		overflow bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "1;2",
			args: args{
				x: 1,
				y: 2,
			},
			want: want{
				result: 3,
			},
		},
		{
			name: "overflow_max",
			args: args{
				x: math.MaxUint - 2,
				y: 3,
			},
			want: want{
				overflow: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := Add(tt.args.x, tt.args.y)
			if !tt.want.overflow {
				if !ok {
					t.Fatalf("unexpected overflow")
				}
				if result != tt.want.result {
					t.Fatalf("%d expected, got %d", tt.want.result, result)
				}
			} else {
				if ok {
					t.Fatalf("overflow expected, got %d", result)
				}
			}
		})
	}
}

func TestSub(t *testing.T) {
	type args struct {
		x uint
		y uint
	}
	type want struct {
		result   uint
		overflow bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "2;1",
			args: args{
				x: 2,
				y: 1,
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "overflow_min",
			args: args{
				x: 2,
				y: 3,
			},
			want: want{
				overflow: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := Sub(tt.args.x, tt.args.y)
			if !tt.want.overflow {
				if !ok {
					t.Fatalf("unexpected overflow")
				}
				if result != tt.want.result {
					t.Fatalf("%d expected, got %d", tt.want.result, result)
				}
			} else {
				if ok {
					t.Fatalf("overflow expected, got %d", result)
				}
			}
		})
	}
}
