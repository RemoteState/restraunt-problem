package main

import (
	"reflect"
	"testing"
)

func TestTopThreeMenuItems(t *testing.T) {
	tests := []struct {
		ordersFilePath string
		expected       []int
		hasError       bool
	}{
		{
			ordersFilePath: "testdata/valid_orders.txt",
			expected:       []int{100, 200, 300},
			hasError:       false,
		},
		{
			ordersFilePath: "testdata/invalid_orders.txt",
			expected:       nil,
			hasError:       true,
		},
		{
			ordersFilePath: "testdata/empty_orders.txt",
			expected:       nil,
			hasError:       true,
		},
	}
	for _, test := range tests {
		got, err := topThreeMenuItems(test.ordersFilePath)
		if test.hasError && err == nil {
			t.Errorf("Expected an error but got nil for orders file path %q", test.ordersFilePath)
		}
		if !test.hasError && err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("Expected %v but got %v", test.expected, got)
		}
	}
}
