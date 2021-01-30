package main

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func Test_saveProduct(t *testing.T) {
	type args struct {
		cnx echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := saveProduct(tt.args.cnx); (err != nil) != tt.wantErr {
				t.Errorf("saveProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
