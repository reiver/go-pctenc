package pctenc

import (
	"testing"
)

func TestUpperHexDig(t *testing.T) {
	tests := []struct{
		HexDig  rune
		Expected rune
	}{
		{
			HexDig:   '0',
			Expected: '0',
		},
		{
			HexDig:   '1',
			Expected: '1',
		},
		{
			HexDig:   '2',
			Expected: '2',
		},
		{
			HexDig:   '3',
			Expected: '3',
		},
		{
			HexDig:   '4',
			Expected: '4',
		},
		{
			HexDig:   '5',
			Expected: '5',
		},
		{
			HexDig:   '6',
			Expected: '6',
		},
		{
			HexDig:   '7',
			Expected: '7',
		},
		{
			HexDig:   '8',
			Expected: '8',
		},
		{
			HexDig:   '9',
			Expected: '9',
		},



		{
			HexDig:   'A',
			Expected: 'A',
		},
		{
			HexDig:   'B',
			Expected: 'B',
		},
		{
			HexDig:   'C',
			Expected: 'C',
		},
		{
			HexDig:   'D',
			Expected: 'D',
		},
		{
			HexDig:   'E',
			Expected: 'E',
		},
		{
			HexDig:   'F',
			Expected: 'F',
		},



		{
			HexDig:   'a',
			Expected: 'A',
		},
		{
			HexDig:   'b',
			Expected: 'B',
		},
		{
			HexDig:   'c',
			Expected: 'C',
		},
		{
			HexDig:   'd',
			Expected: 'D',
		},
		{
			HexDig:   'e',
			Expected: 'E',
		},
		{
			HexDig:   'f',
			Expected: 'F',
		},
	}

	for testNumber, test := range tests {
		actual := upperHexDig(test.HexDig)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual upper hex-dig is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("HEX-DIG:  %q (%U)", test.HexDig, test.HexDig)
			continue
		}
	}
}
