package tempo

import (
	"testing"
)

func TestCheckInterval(t *testing.T) {
	type args struct {
		nower    Nower
		since    Time
		interval Ticks
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
			name: "1;1;1",
			args: args{
				nower: NowerFunc(func() Time {
					return At(1)
				}),
				since:    At(1),
				interval: 1,
			},
			want: want{
				result: true,
			},
		},
		{
			name: "10;1;10",
			args: args{
				nower: NowerFunc(func() Time {
					return At(10)
				}),
				since:    At(1),
				interval: 10,
			},
			want: want{
				result: true,
			},
		},
		{
			name: "9;1;10",
			args: args{
				nower: NowerFunc(func() Time {
					return At(9)
				}),
				since:    At(1),
				interval: 10,
			},
			want: want{
				result: false,
			},
		},
		{
			name: "10;1;9",
			args: args{
				nower: NowerFunc(func() Time {
					return At(10)
				}),
				since:    At(1),
				interval: 9,
			},
			want: want{
				result: false,
			},
		},
		{
			name: "1;10;10",
			args: args{
				nower: NowerFunc(func() Time {
					return At(1)
				}),
				since:    At(10),
				interval: 10,
			},
			want: want{
				result: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CheckInterval(tt.args.nower, tt.args.since, tt.args.interval)
			if result != tt.want.result {
				t.Fatalf("%t expected, got %t", tt.want.result, result)
			}
		})
	}
}
