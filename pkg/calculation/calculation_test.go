package calculation_test

import (
	"CALC_GO/pkg/calculation"
	"testing"
)


func TestCalc(t *testing.T){
		testCasesFail := []struct {
		name        string
		expression  string
		expectedErr error
	}{
		{
			name:       "Not enough operands first",
			expression: "1+1*",
			expectedErr: calculation.ErrNotEnoughOperands,
		},
		{
			name:       "Not enough operands second",
			expression: "2+2**2",
			expectedErr: calculation.ErrNotEnoughOperands,
		},
		{
			name:       "Incorrect placement of brackets",
			expression: "((2+2-*(2",
			expectedErr: calculation.ErrIncorrectBracketPlacement,
		},
		{
			name:       "An empty expression",
			expression: "",
			expectedErr: calculation.ErrEmptyExpression,
		},
	}

	for _, tc := range testCasesFail {
		t.Run(tc.name, func(t *testing.T) {
			val, err := calculation.Calc(tc.expression)
			if err != tc.expectedErr{
				t.Fatalf("expression %s is invalid but result  %f was obtained", tc.expression, val)
			}
		})
	}

	testCaseSuccess := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "Addition",
			expression:     "3 + 4",
			expectedResult: 7.0,
		},
		{
			name:           "Subtraction",
			expression:     "10 - 2",
			expectedResult: 8.0,
		},
		{
			name:           "Multiplication",
			expression:     "2 * 3",
			expectedResult: 6.0,
		},
		{
			name:           "Division",
			expression:     "12 / 3",
			expectedResult: 4.0,
		},
		{
			name:           "MultiExample",
			expression:     "10 * (4 - 5) + 1",
			expectedResult: -9.0,
		},
	}

	for _, tc := range testCaseSuccess {
		t.Run(tc.name, func(t *testing.T) {
			val, err := calculation.Calc(tc.expression)
			if err != nil {
				t.Fatalf("successful case %s returns error", tc.expression)
			}
			if val != tc.expectedResult {
				t.Fatalf("%f should be equal %f", val, tc.expectedResult)
			}
		})
	}
}


