package basic

import "testing"

func Test_Max(t *testing.T) {
	// A really basic test case
	r := Max(1, 20)
	if r != 20 {
		// t.Error(f) is used to trace errors and keep going with the rest of the test function
		t.Errorf("Max(1,20) expecting 20, got %d", r)
	}

	r = Max(30, 20)
	if r != 30 {
		t.Errorf("Max(30,20) expecting 30, got %d", r)
	}

}
