package distance

import (
	"errors"
	"math"

	"github.com/danielstepan/adventofcode/internal/util"
)

func SumOfDifference(a, b []int) (int, error) {
	if !util.AreIntSlicesEqualLength(a, b) {
		return 0, errors.New("slices must be of equal length")
	}

	sum := 0.0
	for i := 0; i < len(a); i++ {
		sum += math.Abs(float64(a[i]) - float64(b[i]))
	}
	return int(sum), nil
}
