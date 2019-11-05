package utils

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestBubbleSortWorstCase(t *testing.T) {
	// Initialization:
	els := []int {9, 8, 7, 6, 5}

	// Execution:
	Sort(els)

	//Validation:
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))

	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}

func generateRandomSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func BenchmarkBubbleSortNative(b *testing.B) {
	els := generateRandomSlice(100)
	for j :=0; j < b.N; j ++ {
		Sort(els)
	}
}

func BenchmarkBubbleSortOwn(b *testing.B) {
	els := generateRandomSlice(100)
	for j :=0; j < b.N; j ++ {
		Sort(els)
	}
}