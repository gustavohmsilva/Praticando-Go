package dataparser

import (
	"reflect"
	"testing"
)

func Test_badReadFile(t *testing.T) {
	p := Product{
		Name:     "Amazon Kindle PaperWhite (Nueva Version)",
		Category: "Dispositivos Amazon",
		Desc: `Te presentamos Kindle: ahora con luz frontal ajustable
para que puedas leer donde y cuando quieras. El Kindle está   
diseñado para la lectura y dispone de una pantalla táctil de
alto contraste que se lee como el papel impreso, sin ningún
reflejo, incluso bajo la luz del sol.`,
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    Product
		wantErr bool
	}{
		{"Test file 001.txt", args{"001.txt"}, p, false},
		{"Test file 002.txt", args{"002.txt"}, Product{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := badReadFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("badReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("badReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadFile(t *testing.T) {
	p := Product{
		Name:     "Amazon Kindle PaperWhite (Nueva Version)",
		Category: "Dispositivos Amazon",
		Desc: `Te presentamos Kindle: ahora con luz frontal ajustable
para que puedas leer donde y cuando quieras. El Kindle está   
diseñado para la lectura y dispone de una pantalla táctil de
alto contraste que se lee como el papel impreso, sin ningún
reflejo, incluso bajo la luz del sol.`,
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    Product
		wantErr bool
	}{
		{"Test file 001.txt", args{"001.txt"}, p, false},
		{"Test file 002.txt", args{"002.txt"}, Product{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
