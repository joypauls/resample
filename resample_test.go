package resample

import "testing"

func TestSimpleSampler(t *testing.T) {
	s := MakeSimpleSampler([]float64{12.2, 54.1, 77.7})

	// test size, w/o replacement
	sample, err := s.Sample(2, false)
	if res, exp := len(sample), 2; res != exp && err != nil {
		t.Errorf("Result: %d, Wanted: %d", res, exp)
	}

	// test size, w/ replacement
	sample, err = s.Sample(2, true)
	if res, exp := len(sample), 2; res != exp && err != nil {
		t.Errorf("Result: %d, Wanted: %d", res, exp)
	}

	// test prevention of oversampling
	sample, err = s.Sample(4, true)
	if err == nil {
		t.Errorf("Result: %e, Wanted: nil", err)
	}
}
