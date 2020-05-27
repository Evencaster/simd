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

func (m *TwoDimensional) Copy() *TwoDimensional {
	duplicate := make([][]float64, len(m.matrix))
	for i := range m.matrix {
		duplicate[i] = make([]float64, len(m.matrix[i]))
		copy(duplicate[i], m.matrix[i])
	}
	return &TwoDimensional{
		matrix: duplicate,
	}
}

func Negative(m *TwoDimensional) *TwoDimensional {
	res := m.Copy()
	for i := range m.matrix {
		for j := range m.matrix[i] {
			res.matrix[i][j] *= -1
		}
	}
	return res
}

func (m *TwoDimensional) MultiplicationInt(n int) {
	for i := range m.matrix {
		for j := range m.matrix[i] {
			m.matrix[i][j] *= float64(n)
		}
	}
}

func MultiplicationInt(first *TwoDimensional, n int) *TwoDimensional {
	m := first.Copy()
	m.MultiplicationInt(n)
	return m
}

func SumInt(first *TwoDimensional, n int) (*TwoDimensional, error) {
	m := first.Copy()
	err := m.SumInt(n)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func Multiplication(first, second *TwoDimensional) (*TwoDimensional, error) {
	if first.N() != second.M() {
		return nil, errors.New("diff size")
	}
	res := NewTwoDimensionalWithValue(first.M(), second.N(), 0)
	for i := range res.matrix {
		for j := range res.matrix[i] {
			for k := range second.matrix {
				res.matrix[i][j] += first.matrix[i][k] * second.matrix[k][j]
			}
		}
	}
	return res, nil
}

func MustMultiplication(first, second *TwoDimensional) *TwoDimensional {
	if first.N() != second.M() {
		panic("diff size")
	}
	res := NewTwoDimensionalWithValue(first.M(), second.N(), 0)
	for i := range res.matrix {
		for j := range res.matrix[i] {
			for k := range second.matrix {
				res.matrix[i][j] += first.matrix[i][k] * second.matrix[k][j]
			}
		}
	}
	return res
}

func MustSum(first, second *TwoDimensional) *TwoDimensional {
	if first.M() != second.M() || first.N() != second.N() {
		panic("diff size")
	}
	res := first.Copy()
	err := res.Sum(second)
	if err != nil {
		panic(err)
	}
	return res
}

func (m *TwoDimensional) Sum(matrix *TwoDimensional) error {
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

func (m *TwoDimensional) SumInt(n int) error {
	matrix := NewTwoDimensionalWithValue(m.M(), m.N(), n)
	return m.Sum(matrix)
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
