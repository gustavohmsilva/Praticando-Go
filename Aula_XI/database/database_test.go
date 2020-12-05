package database

import (
	"testing"

	"github.com/gustavohmsilva/Praticando-Go/Aula_XI/dataparser"
	_ "github.com/mattn/go-sqlite3"
)

func TestInsertProduct(t *testing.T) {
	type args struct {
		p dataparser.Product
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			"Test inserting a product in the database",
			args{
				dataparser.Product{
					Name:     "Mouse PS/2",
					Category: "mouses",
					Desc:     "A basic mouse with a PS/2 connection to the computer",
				},
			},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InsertProduct(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InsertProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
