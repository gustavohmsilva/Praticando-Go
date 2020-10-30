package main

import (
	"testing"
)

func TestMiles2Km(t *testing.T) {
	var got, want float64

	t.Run("testing with 1 mile", func(t *testing.T) {
		got = milesToKm(1.0)
		want = 1.609

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})

	t.Run("testing with 5 miles", func(t *testing.T) {
		got = milesToKm(5.0)
		want = 8.045

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})

	t.Run("testing with 52.6 miles", func(t *testing.T) {
		got = milesToKm(52.6)
		want = 84.6334

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
}
