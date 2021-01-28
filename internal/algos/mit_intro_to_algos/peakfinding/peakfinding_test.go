package peakfinding

import (
	"math/rand"
	"testing"
)

func TestPeakFinding(t *testing.T) {
	input := []int{1, 3, 5, 5, 4, 6, 3, 2, 1, 4, 7, 3, 4, 2, 5}
	pos := Peakfind(input)

	if pos != 3 {
		t.Errorf("Peak finding failed, expected: %v; got: %v", 3, pos)
	}
}

func TestPeakFindingBrute(t *testing.T) {
	input := []int{1, 3, 5, 6, 4, 6, 3, 2, 1, 4, 7, 3, 4, 2, 5}
	pos, _ := PeakFindBruteForce(input)

	if pos != 3 {
		t.Errorf("Peak finding failed, expected: %v; got: %v", 3, pos)
	}
}

func BenchmarkPeakFinding(b *testing.B) {
	for i := 0; i < 100; i++ {
		b.StopTimer()
		input := rand.Perm(10)
		b.StartTimer()
		Peakfind(input)
	}
}

func BenchmarkPeakFindingBrute(b *testing.B) {
	for i := 0; i < 100; i++ {
		b.StopTimer()
		input := rand.Perm(10)
		b.StartTimer()
		PeakFindBruteForce(input)
	}
}
