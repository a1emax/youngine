package clock

import (
	"testing"
)

func TestAt(t *testing.T) {
	type args struct {
		ticks Ticks
	}
	type want struct {
		result Time
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "0",
			args: args{
				ticks: 0,
			},
			want: want{
				result: Time{0},
			},
		},
		{
			name: "-1",
			args: args{
				ticks: -1,
			},
			want: want{
				result: Time{0},
			},
		},
		{
			name: "1",
			args: args{
				ticks: 1,
			},
			want: want{
				result: Time{1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := At(tt.args.ticks)
			if result != tt.want.result {
				t.Fatalf("(%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestTime_Ticks(t *testing.T) {
	type args struct {
		t Time
	}
	type want struct {
		result Ticks
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "empty_literal",
			args: args{
				t: Time{},
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "0",
			args: args{
				t: Time{0},
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "1",
			args: args{
				t: Time{1},
			},
			want: want{
				result: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.t.Ticks()
			if result != tt.want.result {
				t.Fatalf("(%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestTime_IsZero(t *testing.T) {
	type args struct {
		t Time
	}
	type want struct {
		result bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "empty_literal",
			args: args{
				t: Time{},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "0",
			args: args{
				t: Time{0},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "1",
			args: args{
				t: Time{1},
			},
			want: want{
				result: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.t.IsZero()
			if result != tt.want.result {
				t.Fatalf("(%t expected, got %t", tt.want.result, result)
			}
		})
	}
}

func TestTime_Add(t *testing.T) {
	type args struct {
		t Time
		d Ticks
	}
	type want struct {
		result Time
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "1;2",
			args: args{
				t: Time{1},
				d: 2,
			},
			want: want{
				result: Time{3},
			},
		},
		{
			name: "1;-2",
			args: args{
				t: Time{1},
				d: -2,
			},
			want: want{
				result: Time{0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.t.Add(tt.args.d)
			if result != tt.want.result {
				t.Fatalf("(%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestTime_Sub(t *testing.T) {
	type args struct {
		t Time
		u Time
	}
	type want struct {
		result Ticks
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "0;0",
			args: args{
				t: Time{0},
				u: Time{0},
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "0;1",
			args: args{
				t: Time{0},
				u: Time{1},
			},
			want: want{
				result: -1,
			},
		},
		{
			name: "1;0",
			args: args{
				t: Time{1},
				u: Time{0},
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "1;1",
			args: args{
				t: Time{1},
				u: Time{1},
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "2;1",
			args: args{
				t: Time{2},
				u: Time{1},
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "1;2",
			args: args{
				t: Time{1},
				u: Time{2},
			},
			want: want{
				result: -1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.t.Sub(tt.args.u)
			if result != tt.want.result {
				t.Fatalf("(%d expected, got %d", tt.want.result, result)
			}
		})
	}
}
