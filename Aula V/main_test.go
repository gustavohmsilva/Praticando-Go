package main

import (
	"reflect"
	"testing"
)

func Test_parseData(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{
			"Test I",
			args{[]string{"60,1", "58,3", "82,1", "62,8", "75,6"}},
			[]float64{60.1, 58.3, 82.1, 62.8, 75.6},
			false,
		},
		{
			"Test II",
			args{[]string{"45,8", "55,6", "68,4", "50,0", "65,2", "47,3", "48,6", "50,2"}},
			[]float64{45.8, 55.6, 68.4, 50.0, 65.2, 47.3, 48.6, 50.2},
			false,
		},
		{
			"Test III",
			args{[]string{"60,1", "58,3", "82,1", "62;8", "75,6"}},
			[]float64{60.1, 58.3, 82.1},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseData(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateMedian(t *testing.T) {
	type args struct {
		dat []float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			"Result I",
			args{[]float64{60.1, 58.3, 82.1, 62.8, 75.6}},
			67.78,
			false,
		},
		{
			"Result Ii",
			args{[]float64{45.8, 55.6, 68.4, 50.0, 65.2, 47.3, 48.6, 50.2}},
			53.8875,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateMedian(tt.args.dat)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateMedian() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculateMedian() = %v, want %v", got, tt.want)
			}
		})
	}
}
