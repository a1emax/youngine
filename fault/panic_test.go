package fault

import (
	"errors"
	"testing"
)

func TestRecover__Panic(t *testing.T) {
	testErr := errors.New("error")

	err := Recover(func() {
		panic(testErr)
	})
	if err == nil {
		t.Fatalf("error expected, got nil")
	}
	if !errors.Is(err, testErr) {
		t.Fatalf("unexpected error: %+v", err)
	}
}

func TestRecover__NoPanic(t *testing.T) {
	err := Recover(func() {})
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
}

func TestRecovered(t *testing.T) {
	type args struct {
		v any
	}
	type want struct {
		errString string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "nil",
			args: args{
				v: nil,
			},
			want: want{
				errString: "",
			},
		},
		{
			name: "int",
			args: args{
				v: 1,
			},
			want: want{
				errString: "panic recovered: 1",
			},
		},
		{
			name: "error",
			args: args{
				v: errors.New("error"),
			},
			want: want{
				errString: "error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Recovered(tt.args.v)
			if tt.want.errString == "" {
				if err != nil {
					t.Fatalf("unexpected error: %+v", err)
				}
			} else {
				if err == nil {
					t.Fatalf("error expected, got nil")
				}
				if err.Error() != tt.want.errString {
					t.Fatalf("unexpected error: %+v", err)
				}
			}
		})
	}
}
