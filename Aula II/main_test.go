package main

import "testing"

func Test_countCoins(t *testing.T) {
	type args struct {
		c coins
	}
	tests := []struct {
		name      string
		args      args
		wantEuro  int
		wantCents int
	}{
		{"Test 10 1 cents coins", args{coins{10, 0, 0, 0, 0, 0, 0, 0}}, 0, 10},
		{"Test 10 2 cents coins", args{coins{0, 10, 0, 0, 0, 0, 0, 0}}, 0, 20},
		{"Test 10 5 cents coins", args{coins{0, 0, 10, 0, 0, 0, 0, 0}}, 0, 50},
		{"Test 10 10 cents coins", args{coins{0, 0, 0, 10, 0, 0, 0, 0}}, 1, 0},
		{"Test 10 20 cents coins", args{coins{0, 0, 0, 0, 10, 0, 0, 0}}, 2, 0},
		{"Test 10 50 cents coins", args{coins{0, 0, 0, 0, 0, 10, 0, 0}}, 5, 0},
		{"Test 10 1 euro coins", args{coins{0, 0, 0, 0, 0, 0, 10, 0}}, 10, 0},
		{"Test 10 2 euros coins", args{coins{0, 0, 0, 0, 0, 0, 0, 10}}, 20, 0},
		{"Test 10 of each coin", args{coins{10, 10, 10, 10, 10, 10, 10, 10}}, 38, 80},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := countCoins(tt.args.c)
			if got != tt.wantEuro {
				t.Errorf("countCoins() got = %v, want %v", got, tt.wantEuro)
			}
			if got1 != tt.wantCents {
				t.Errorf("countCoins() got1 = %v, want %v", got1, tt.wantCents)
			}
		})
	}
}
