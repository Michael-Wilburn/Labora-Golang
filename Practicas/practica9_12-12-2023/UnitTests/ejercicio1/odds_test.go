package main

import "testing"

func TestAdd(t *testing.T) {

	type Test struct {
		n           int
		expectedSum int
	}

	test := []Test{
		{
			n:           0,
			expectedSum: 0,
		},
		{
			n:           1,
			expectedSum: 1,
		},
		{
			n:           2,
			expectedSum: 4,
		},
		{
			n:           5,
			expectedSum: 25,
		},
	}

	for _, test := range test {
		computedSum := sumarPrimerosImpares(test.n)
		if computedSum != test.expectedSum {
			t.Errorf("suma de Impares(%d) = %d, esperado %d", test.n, computedSum, test.expectedSum)
		}
	}

}
