package main

import "testing"

func TestTLSSupport(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			"abba outside square brackets",
			"abba[mnop]qrst",
			true,
		},
		{
			"bddb is within square brackets, even though xyyx is outside square brackets",
			"abcd[bddb]xyyx",
			false,
		},
		{
			"aaaa is invalid; the interior characters must be different",
			"aaaa[qwer]tyui",
			false,
		},
		{
			"oxxo is outside square brackets, even though it's within a larger string",
			"ioxxoj[asdfgh]zxcvbn",
			true,
		},
	}

	for _, tt := range tests {
		if got := checkTLSSupport(tt.input); got != tt.want {
			t.Errorf("%q. checkTLSSupport(%s) = %v, want %v", tt.name, tt.input, got, tt.want)
		}
	}
}
