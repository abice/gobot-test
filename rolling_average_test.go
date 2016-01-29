// rolling_average_test.go
package main

import (
	"testing"
)

func TestAverageCompute(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	averager := NewAverager("test", 10)

	t.Log("Computing initial average")
	avg := averager.Compute()

	if avg != 0 {
		t.Logf("Initial average should be 0")
		t.Fail()
	} else {
		t.Logf("Average: %d", avg)
	}

	averager.Add(1)

	avg = averager.Compute()

	if avg != 1 {
		t.Logf("Failed computing Average: %d", avg)
		t.Fail()
	} else {
		t.Logf("Average: %d", avg)
	}

	for i := 2; i < 20; i++ {
		averager.Add(i)
	}

	avg = averager.Compute()

	if avg <= 10 {
		t.Logf("Failed computing Average: %d", avg)
		t.Fail()
	} else {
		t.Logf("Average: %d", avg)
	}

}
