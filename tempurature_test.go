package uconv

import "testing"

func TestCelsiusToFahrenheit(t *testing.T) {
	cases := []struct {
		in   Celsius
		want Fahrenheit
	}{
		{AbsoluteZeroC, AbsoluteZeroF},
		{FreezingC, FreezingF},
		{BoilingC, BoilingF},
	}
	for _, c := range cases {
		got := CelsiusToFahrenheit(c.in)
		if got != c.want {
			t.Errorf("CelsiusToFahrenheit(%f) == %f, want %f", c.in, got, c.want)
		}
	}
}
