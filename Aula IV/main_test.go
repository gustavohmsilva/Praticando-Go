package main

import "testing"

func Test_shippingPrice(t *testing.T) {
	type args struct {
		w      washers
		weight float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Data given by client",
			args{washers{0.2, 2.0, 1.0, 1000, 2.50}, 7.8},
			"R$9.19",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shippingPrice(tt.args.w, tt.args.weight); got != tt.want {
				t.Errorf("shippingPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
