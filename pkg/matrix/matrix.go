package matrix

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// TwoDimensional defines two dimensional matrix.
type TwoDimensional struct {
	matrix [][]float64
}

// NewTwoDimensional returns new TwoDimensional matrix with empty data.
func NewTwoDimensional() *TwoDimensional {
	return &TwoDimensional{
		matrix: [][]float64{},
	}
}

// NewTwoDimensionalGenerated returns generated TwoDimensional matrix with m*n size with values [-1;1].
func NewTwoDimensionalGenerated(m, n int) *TwoDimensional {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	matrix := make([][]float64, m)
	for i := 0; i < m; i++ {
		subSlice := make([]float64, n)
		matrix[i] = subSlice
		for j := 0; j < n; j++ {
			matrix[i][j] = -1 + r1.Float64()*2
		}
	}
	return &TwoDimensional{
		matrix: matrix,
	}
}

func NewTwoDimensionalWithValue(m, n, v int) *TwoDimensional {
	matrix := make([][]float64, m)
	for i := 0; i < m; i++ {
		subSlice := make([]float64, n)
		matrix[i] = subSlice
		for j := 0; j < n; j++ {
			matrix[i][j] = float64(v)
		}
	}
	return &TwoDimensional{
		matrix: matrix,
	}
}

func (m *TwoDimensional) M() int {
	return len(m.matrix)
}

func (m *TwoDimensional) N() int {
	if m.M() == 0 {
		return 0
	}
	return len(m.matrix[0])
}

func (m *TwoDimensional) Multiplication(n int) {
	for i := range m.matrix {
		for j := range m.matrix[i] {
			m.matrix[i][j] *= float64(n)
		}
	}
}

func (m *TwoDimensional) Sum(matrix TwoDimensional) error {
	if m.M() != matrix.M() || m.N() != matrix.N() {
		return errors.New("diff size")
	}
	for i := 0; i < m.M(); i++ {
		for j := 0; j < m.N(); j++ {
			m.matrix[i][j] += matrix.matrix[i][j]
		}
	}
	return nil
}

func (m *TwoDimensional) String() string {
	b := &strings.Builder{}
	for i := range m.matrix {
		for _, v := range m.matrix[i] {
			b.WriteString(strconv.FormatFloat(v, 'f', 2, 64))
			b.WriteString(" ")
		}
		b.WriteString("\n")
	}
	return b.String()
}
