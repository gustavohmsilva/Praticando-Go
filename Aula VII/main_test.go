package main

import "testing"

func Test_reverseMessage(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"O rato roeu...",
			args{"O Rato Roeu a Roupa do Rei de Roma"},
			"amoR ed ieR od apuoR a ueoR otaR O",
			false,
		},
		{
			"Descrição Ikea...",
			args{"Aprenda a montar uma cozinha desde o zero. Antes de começar é imprescindível planejar bem o que quer..."},
			"...reuq euq o meb rajenalp levídnicserpmi é raçemoc ed setnA .orez o edsed ahnizoc amu ratnom a adnerpA",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := reverseMessage(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("reverseMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("reverseMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
