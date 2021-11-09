package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloor(t *testing.T) {
	tests := map[string]struct {
		input         string
		expectedFloor int
	}{
		"floor0a": {
			input:         "(())",
			expectedFloor: 0,
		},
		"floor0b": {
			input:         "()()",
			expectedFloor: 0,
		},
		"floor3a": {
			input:         "(((",
			expectedFloor: 3,
		},
		"floor3b": {
			input:         "(()(()(",
			expectedFloor: 3,
		},
		"floor3c": {
			input:         "))(((((",
			expectedFloor: 3,
		},
		"floor-1a": {
			input:         "())",
			expectedFloor: -1,
		},
		"floor-1b": {
			input:         "))(",
			expectedFloor: -1,
		},
		"floor-3a": {
			input:         ")))",
			expectedFloor: -3,
		},
		"floor-3b": {
			input:         ")())())",
			expectedFloor: -3,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			floor := floor(test.input)
			assert.Equal(t, test.expectedFloor, floor)
		})
	}
}

func TestMinus1Floor(t *testing.T) {
	tests := map[string]struct {
		input       string
		expectedPos int
	}{
		"floor0a": {
			input:       "(()))",
			expectedPos: 5,
		},
		"floor0b": {
			input:       "()()",
			expectedPos: -1,
		},
		"floor3a": {
			input:       "(((",
			expectedPos: -1,
		},
		"floor3b": {
			input:       "(()(()(",
			expectedPos: -1,
		},
		"floor3c": {
			input:       "))(((((",
			expectedPos: 1,
		},
		"floor-1a": {
			input:       "())",
			expectedPos: 3,
		},
		"floor-1b": {
			input:       "))(",
			expectedPos: 1,
		},
		"floor-3a": {
			input:       ")))",
			expectedPos: 1,
		},
		"floor-3b": {
			input:       ")())())",
			expectedPos: 1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			pos := minus1Floor(test.input)
			assert.Equal(t, test.expectedPos, pos)
		})
	}
}
