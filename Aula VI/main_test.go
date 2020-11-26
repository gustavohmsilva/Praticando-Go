package main

import (
	"testing"
)

func Test_generateFizzBuzz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"working 10 entry",
			args{10},
			"1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz",
			false,
		},
		{
			"working 15 entry",
			args{15},
			"1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizzbuzz",
			false,
		},
		{
			"Working 1 entry",
			args{1},
			"1",
			false,
		},
		{
			"Non-working 0 entry",
			args{0},
			"",
			true,
		},
		{
			"Non-working -10 entry",
			args{-10},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateFizzBuzz(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateFizzBuzz() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("generateFizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataEntry(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Working Example", args{"10"}, 10, false},
		{"Non-Working Example", args{"dez"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dataEntry(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("dataEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("dataEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}
