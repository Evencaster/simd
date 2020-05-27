package simd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/illfate2/simd/pkg/matrix"
)

func TestProcessor_Result(t *testing.T) {
	const p = 2
	const m = 3
	const q = 2
	args := ProcessorArgs{
		A: matrix.NewTwoDimensionalGenerated(p, m),
		B: matrix.NewTwoDimensionalGenerated(m, q),
		E: matrix.NewTwoDimensionalGenerated(1, m),
		G: matrix.NewTwoDimensionalGenerated(p, q),
		K: 5,
	}
	processor := NewProcessor(args)
	got, err := processor.Result()
	require.NoError(t, err)
	fmt.Print(got)
}
