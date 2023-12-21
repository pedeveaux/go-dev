package goalpacegenerator

import "testing"

func TestConvertPace(t *testing.T) {
	testCases := []struct {
		name     string
		pacemeters int
		goalDistance int
		goalMinutes  int
		goalSeconds  int
		expected paceUnit
	}{
		{
			name:     "Test case 1",
			pacemeters: 400,
			goalDistance: 1600,
			goalMinutes:  4,
			goalSeconds:  40,
			expected: paceUnit{meters: 400, minutes: 0, seconds: 70},
		},
		{
			name:     "Test case 2",
			pacemeters: 200,
			goalDistance: 1600,
			goalMinutes:  5,
			goalSeconds:  0,
			expected: paceUnit{meters: 200, minutes: 0, seconds: 37},
		},
		{
			name:     "Test case 3",
			pacemeters: 1000,
			goalDistance: 1600,
			goalMinutes:  5,
			goalSeconds:  0,
			expected: paceUnit{meters: 1000, minutes: 3, seconds: 7},
		},
		{
			name:     "Test case 4",
			pacemeters: 1200,
			goalDistance: 1600,
			goalMinutes:  5,
			goalSeconds:  0,
			expected: paceUnit{meters: 1200, minutes: 3, seconds: 45},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			myInstance := paceUnit{meters: tc.pacemeters}
			myInstance.convertPace(tc.goalMinutes, tc.goalSeconds, tc.goalDistance)
			if myInstance != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, myInstance)
			}
		})
	}
}
