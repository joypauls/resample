package resample

import "testing"

func TestSimpleSampler(t *testing.T) {
	s := SimpleSampler{data: []float64{12.2, 54.1, 77.7}}

	// test size
	if res, exp := len(s.Sample(1, false)), 1; res != exp {
		t.Errorf("Result: %d, Wanted: %d", res, exp)
	}
}
