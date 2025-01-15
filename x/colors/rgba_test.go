package colors

import (
	"testing"
)

func TestRGBA_RGBA(t *testing.T) {
	type args struct {
		c RGBA
	}
	type want struct {
		r uint32
		g uint32
		b uint32
		a uint32
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "0x4D99FF80",
			args: args{
				c: RGBA{0x4D, 0x99, 0xFF, 0x80},
			},
			want: want{
				r: 0x4D4D,
				g: 0x9999,
				b: 0xFFFF,
				a: 0x8080,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b, a := tt.args.c.RGBA()
			if r != tt.want.r || g != tt.want.g || b != tt.want.b || a != tt.want.a {
				t.Fatalf("(%d, %d, %d, %d) expected, got (%d, %d, %d, %d)",
					tt.want.r, tt.want.g, tt.want.b, tt.want.a,
					r, g, b, a,
				)
			}
		})
	}
}

func TestRGBA_R(t *testing.T) {
	type args struct {
		c RGBA
	}
	type want struct {
		result uint8
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "0x4D99FF80",
			args: args{
				c: RGBA{0x4D, 0x99, 0xFF, 0x80},
			},
			want: want{
				result: 0x4D,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.c.R()
			if result != tt.want.result {
				t.Fatalf("(%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestRGBA_G(t *testing.T) {
	type args struct {
		c RGBA
	}
	type want struct {
		result uint8
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "0x4D99FF80",
			args: args{
				c: RGBA{0x4D, 0x99, 0xFF, 0x80},
			},
			want: want{
				result: 0x99,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.c.G()
			if result != tt.want.result {
				t.Fatalf("(%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestRGBA_B(t *testing.T) {
	type args struct {
		c RGBA
	}
	type want struct {
		result uint8
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "0x4D99FF80",
			args: args{
				c: RGBA{0x4D, 0x99, 0xFF, 0x80},
			},
			want: want{
				result: 0xFF,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.c.B()
			if result != tt.want.result {
				t.Fatalf("(%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestRGBA_A(t *testing.T) {
	type args struct {
		c RGBA
	}
	type want struct {
		result uint8
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "0x4D99FF80",
			args: args{
				c: RGBA{0x4D, 0x99, 0xFF, 0x80},
			},
			want: want{
				result: 0x80,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.c.A()
			if result != tt.want.result {
				t.Fatalf("(%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestRGBA_Premul(t *testing.T) {
	type args struct {
		c RGBA
	}
	type want struct {
		result RGBA
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "0x4D99FF80",
			args: args{
				c: RGBA{0x4D, 0x99, 0xFF, 0x80},
			},
			want: want{
				result: RGBA{0x26, 0x4C, 0x80, 0x80},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.c.Premul()
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestRGBA_ToFRGBA(t *testing.T) {
	type args struct {
		c RGBA
	}
	type want struct {
		result FRGBA
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "0x4D99FF80",
			args: args{
				c: RGBA{0x4D, 0x99, 0xFF, 0x80},
			},
			want: want{
				result: FRGBA{0.3, 0.6, 1.0, 0.5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.c.ToFRGBA().Round(1)
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}
