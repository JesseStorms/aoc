package main

import (
	"reflect"
	"testing"
)

func TestSeedToSoil(t *testing.T) {
	route := Route{DestStart: 50, SourceStart: 98, Length: 2, Type: Soil}

	testCases := []struct {
		name     string
		initial  Step
		route    Route
		expected Step
	}{
		{name: "Normal case 1", initial: Step{Value: 98, Type: Seed}, route: route, expected: Step{Value: 50, Type: Soil}},
		{name: "Normal case 2", initial: Step{Value: 99, Type: Seed}, route: route, expected: Step{Value: 51, Type: Soil}},
		{name: "Case when the value is too high", initial: Step{Value: 100, Type: Seed}, route: route, expected: Step{Value: 100, Type: Soil}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := route.Solve(tc.initial)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestSeedToFertilizer(t *testing.T) {
	soilRoute := Route{DestStart: 52, SourceStart: 50, Length: 2, Type: Soil}
	fertilizerRoute := Route{DestStart: 1, SourceStart: 52, Length: 2, Type: Fertilizer}
	testCases := []struct {
		name     string
		initial  Step
		expected Step
	}{
		{name: "Normal case 1", initial: Step{Value: 50, Type: Seed}, expected: Step{Value: 1, Type: Fertilizer}},
		{name: "Normal case 2", initial: Step{Value: 51, Type: Seed}, expected: Step{Value: 2, Type: Fertilizer}},
		{name: "Case when the value is too high", initial: Step{Value: 100, Type: Seed}, expected: Step{Value: 100, Type: Fertilizer}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := fertilizerRoute.Solve(soilRoute.Solve(tc.initial))
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
