package main

import (
	"testing"
)

func TestSolveDay11(t *testing.T) {
	test_string := []string{
		"1abc2\n",
		"pqr3stu8vwx\n",
		"a1b2c3d4e5f\n",
		"treb7uchet",
	}
	test_data := []byte{}
	for _, v := range test_string {
		test_data = append(test_data, v...)
	}
	result := SolveDay11(test_data)
	if result != 142 {
		t.Fatalf("Expected 142, got %v instead", result)
	}
}
