package resample

import (
	"math/rand"
	"time"
)

// Core interface that any sampling method should implement
type Sampler interface {
	// Shuffle() []float64
	Sample(n int, replacement bool) []float64
}

type SimpleSampler struct {
	data []float64
}

func (s *SimpleSampler) Sample(n int, replacement bool) []float64 {
	rand.Seed(time.Now().UnixNano())
	sample := make([]float64, n)
	for i := range sample {
		sample[i] = s.data[rand.Intn(len(s.data))]
	}
	return sample
}
