// rolling_average.go
package main

import (
	"fmt"
)

type average struct {
	name        string
	sampleCount int
	buffer      []int
	index       int
}

type Averager interface {
	Add(value int)
	Compute() (average int)
}

type Average struct {
	Averager
}

func NewAverager(name string, numSamples int) *Average {
	avg := &Average{
		Averager: &average{
			name:        name,
			sampleCount: numSamples,
			index:       0,
			buffer:      make([]int, 0, numSamples),
		},
	}

	return avg
}

func (a *average) Add(value int) {
	if len(a.buffer) < a.sampleCount {
		a.buffer = append(a.buffer, value)
		a.index++
	} else {
		if a.index >= a.sampleCount-1 {
			a.index = 0
		} else {
			a.index++
		}
		a.buffer[a.index] = value
	}

}

func (a *average) Compute() (average int) {
	sum := 0
	buffer_len := len(a.buffer)
	if buffer_len == 0 {
		return 0
	}

	count := 0
	if buffer_len < a.sampleCount {
		count = buffer_len
	} else {
		count = a.sampleCount
	}

	for i := 0; i < count; i++ {
		sum += a.buffer[i]
	}

	fmt.Println("Current length: ", len(a.buffer))
	if sum == 0 {
		return 0
	}
	avg := sum / len(a.buffer)
	return avg
}
