package sample_test

import (
	"testing"
	"webapp/app/sample"
)

func TestHi(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Hi Carlos",
			args: args{name: "Carlos"},
			want: "Hi, Carlos",
		},
		{
			name: "Test Hi Oscar",
			args: args{name: "Oscar"},
			want: "Hi, Oscar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sample.Hi(tt.args.name); got != tt.want {
				t.Errorf("Hi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 2 + 3",
			args: args{a: 2, b: 3},
			want: 5,
		},
		{
			name: "Test -2 + 3",
			args: args{a: -2, b: 3},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sample.Sum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
