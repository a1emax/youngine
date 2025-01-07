package colors

import (
	"testing"

	"github.com/a1emax/youngine/basic"
)

func TestFRGBA_RGBA(t *testing.T) {
	type args struct {
		c FRGBA
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
			name: "rgba(0.3,0.6,1.0,0.5)",
			args: args{
				c: FRGBA{0.3, 0.6, 1.0, 0.5},
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

func TestFRGBA_R(t *testing.T) {
	type args struct {
		c FRGBA
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
			name: "rgba(0.3,0.6,1.0,0.5)",
			args: args{
				c: FRGBA{0.3, 0.6, 1.0, 0.5},
			},
			want: want{
				result: 0.3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.c.R()
			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestFRGBA_G(t *testing.T) {
	type args struct {
		c FRGBA
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
			name: "rgba(0.3,0.6,1.0,0.5)",
			args: args{
				c: FRGBA{0.3, 0.6, 1.0, 0.5},
			},
			want: want{
				result: 0.6,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.c.G()
			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestFRGBA_B(t *testing.T) {
	type args struct {
		c FRGBA
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
			name: "rgba(0.3,0.6,1.0,0.5)",
			args: args{
				c: FRGBA{0.3, 0.6, 1.0, 0.5},
			},
			want: want{
				result: 1.0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.c.B()
			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestFRGBA_A(t *testing.T) {
	type args struct {
		c FRGBA
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
			name: "rgba(0.3,0.6,1.0,0.5)",
			args: args{
				c: FRGBA{0.3, 0.6, 1.0, 0.5},
			},
			want: want{
				result: 0.5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.c.A()
			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestFRGBA_Strict(t *testing.T) {
	type args struct {
		c FRGBA
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
			name: "rgba(-0.3,1.6,1.0,0.5)",
			args: args{
				c: FRGBA{-0.3, 1.6, 1.0, 0.5},
			},
			want: want{
				result: FRGBA{0.0, 1.0, 1.0, 0.5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.c.Strict()
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestFRGBA_Premul(t *testing.T) {
	type args struct {
		c FRGBA
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
			name: "rgba(0.3,0.6,1.0,0.5)",
			args: args{
				c: FRGBA{0.3, 0.6, 1.0, 0.5},
			},
			want: want{
				result: FRGBA{0.15, 0.3, 0.5, 0.5},
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

func TestFRGBA_Round(t *testing.T) {
	type args struct {
		c FRGBA
		n int
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
			name: "rgba(-0.345,1.678,1.0,0.5);1",
			args: args{
				c: FRGBA{-0.345, 1.678, 1.0, 0.5},
				n: 1,
			},
			want: want{
				result: FRGBA{-0.3, 1.7, 1.0, 0.5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.c.Round(tt.args.n)
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestFRGBA_ToRGBA(t *testing.T) {
	type args struct {
		c FRGBA
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
			name: "rgba(0.3,0.6,1.0,0.5)",
			args: args{
				c: FRGBA{0.3, 0.6, 1.0, 0.5},
			},
			want: want{
				result: RGBA{0x4D, 0x99, 0xFF, 0x80},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.c.ToRGBA()
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}
