package colors

import (
	"testing"

	"github.com/a1emax/youngine/basic"
)

func TestFHSLA_RGBA(t *testing.T) {
	type args struct {
		c FHSLA
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
			name: "(214.29,1.0,0.65,0.5)",
			args: args{
				c: FHSLA{214.29, 1.0, 0.65, 0.5},
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

func TestFHSLA_H(t *testing.T) {
	type args struct {
		c FHSLA
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
			name: "(214.29,1.0,0.65,0.5)",
			args: args{
				c: FHSLA{214.29, 1.0, 0.65, 0.5},
			},
			want: want{
				result: 214.29,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.c.H()
			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestFHSLA_S(t *testing.T) {
	type args struct {
		c FHSLA
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
			name: "(214.29,1.0,0.65,0.5)",
			args: args{
				c: FHSLA{214.29, 1.0, 0.65, 0.5},
			},
			want: want{
				result: 1.0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.c.S()
			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestFHSLA_L(t *testing.T) {
	type args struct {
		c FHSLA
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
			name: "(214.29,1.0,0.65,0.5)",
			args: args{
				c: FHSLA{214.29, 1.0, 0.65, 0.5},
			},
			want: want{
				result: 0.65,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.c.L()
			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestFHSLA_A(t *testing.T) {
	type args struct {
		c FHSLA
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
			name: "(214.29,1.0,0.65,0.5)",
			args: args{
				c: FHSLA{214.29, 1.0, 0.65, 0.5},
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

func TestFHSLA_Strict(t *testing.T) {
	type args struct {
		c FHSLA
	}
	type want struct {
		result FHSLA
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(-214.29,1.0,1.65,0.5)",
			args: args{
				c: FHSLA{-214.29, 1.0, 1.65, 0.5},
			},
			want: want{
				result: FHSLA{0.0, 1.0, 1.0, 0.5},
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

func TestFHSLA_Round(t *testing.T) {
	type args struct {
		c FHSLA
		n int
	}
	type want struct {
		result FHSLA
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(-214.29,1.0,1.65,0.5)",
			args: args{
				c: FHSLA{-214.29, 1.0, 1.65, 0.5},
				n: 1,
			},
			want: want{
				result: FHSLA{-214.3, 1.0, 1.7, 0.5},
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

func TestFHSLA_ToFRGBA(t *testing.T) {
	type args struct {
		c FHSLA
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
			name: "(214.29,1.0,0.65,0.5)",
			args: args{
				c: FHSLA{214.29, 1.0, 0.65, 0.5},
			},
			want: want{
				result: RGBA{0x4D, 0x99, 0xFF, 0x80},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.c.ToFRGBA().ToRGBA()
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}
