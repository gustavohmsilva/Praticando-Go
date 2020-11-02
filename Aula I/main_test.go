package main

import "testing"

func Test_milesToKm(t *testing.T) {
	type args struct {
		m float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{"1 mile Test", args{1.0}, 1.609000},
		{"5 miles Test", args{5.0}, 8.045000},
		{"52.6 miles Test", args{52.6}, 84.63339},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := milesToKm(tt.args.m); got != tt.want {
				t.Errorf("milesToKm() = %v, want %v", got, tt.want)
			}
		})
	}
}
