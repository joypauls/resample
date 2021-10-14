package resample

import (
	"errors"
	"math/rand"
	"time"
)

// Core interface that any sampling method should implement
type Sampler interface {
	// Shuffle() []float64
	Sample(n int, replacement bool) ([]float64, error)
}

// is there actually any benefit for a sampler to have data? for resampling probably, for a single sample? meh
type SimpleSampler struct {
	data []float64
}

func (s *SimpleSampler) Sample(k int, replacement bool) ([]float64, error) {
	// disallow oversampling for simple use case
	if k > len(s.data) {
		return []float64{}, errors.New("Invalid sample size requested")
	}
	rand.Seed(time.Now().UnixNano()) // should this always be done?
	sample := make([]float64, k)
	if replacement {
		// with replacement
		for i := 0; i < k; i++ {
			sample[i] = s.data[rand.Intn(len(s.data))]
		}
		return sample, nil
	}
	// without replacement
	idx := rand.Perm(len(s.data))
	for i := 0; i < k; i++ {
		sample[i] = s.data[idx[i]]
	}
	return sample, nil
}

func MakeSimpleSampler(data []float64) SimpleSampler {
	return SimpleSampler{data: data}
}
