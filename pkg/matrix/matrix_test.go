package matrix

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTwoDimensional_String(t *testing.T) {
	type fields struct {
		matrix [][]float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "two row",
			fields: fields{
				[][]float64{
					{1.0, 0.1, -0.1},
					{-1, -0.6, 0.5},
				},
			},
			want: "1.00 0.10 -0.10 \n-1.00 -0.60 0.50 \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &TwoDimensional{
				matrix: tt.fields.matrix,
			}
			res := m.String()
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestTwoDimensional_Multiplication(t *testing.T) {
	src := [][]float64{
		{1.0, 0.1, -0.1},
		{-1, -0.6, 0.5},
	}
	n := 3
	m := &TwoDimensional{
		matrix: src,
	}
	m.MultiplicationInt(n)

	exp := [][]float64{
		{3.0, 0.30000000000000004, -0.30000000000000004},
		{-3, -1.7999999999999998, 1.5},
	}
	assert.Equal(t, exp, m.matrix)
}

func TestTwoDimensional_Sum(t *testing.T) {
	type fields struct {
		matrix [][]float64
	}
	type args struct {
		matrix *TwoDimensional
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expMatrix TwoDimensional
	}{
		{
			name: "",
			fields: fields{
				matrix: [][]float64{
					{0, 0.5, -0.5},
					{-0.3, 0.7, 0.6},
				},
			},
			args: args{
				matrix: &TwoDimensional{matrix: [][]float64{
					{0.2, 0.3, -0.1},
					{0.9, -0.6, -0.2},
				}},
			},
			expMatrix: TwoDimensional{
				matrix: [][]float64{
					{0.2, 0.8, -0.6},
					{0.6, 0.1, 0.4},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &TwoDimensional{
				matrix: tt.fields.matrix,
			}
			err := m.Sum(tt.args.matrix)
			require.NoError(t, err)
			assertMatrix(t, tt.expMatrix, *m)
		})
	}
}

func assertMatrix(t *testing.T, want, got TwoDimensional) {
	require.Equal(t, want.M(), got.M())
	require.Equal(t, want.N(), got.N())
	for i := range want.matrix {
		for j := range want.matrix[i] {
			assert.Equal(t, math.Round(want.matrix[i][j]*100)/100, math.Round(got.matrix[i][j]*100)/100)
		}
	}

}

func TestMultiplication(t *testing.T) {
	type args struct {
		first  *TwoDimensional
		second *TwoDimensional
	}
	tests := []struct {
		name string
		args args
		want *TwoDimensional
	}{
		{
			name: "",
			args: args{
				first: &TwoDimensional{matrix: [][]float64{
					{-2, 1},
					{5, 4},
				}},
				second: &TwoDimensional{matrix: [][]float64{
					{3},
					{-1},
				}},
			},
			want: &TwoDimensional{matrix: [][]float64{
				{-7},
				{11},
			}},
		},
		{
			name: "",
			args: args{
				first: &TwoDimensional{matrix: [][]float64{
					{2, -3},
					{4, -6},
				}},
				second: &TwoDimensional{matrix: [][]float64{
					{9, -6},
					{6, -4},
				}},
			},
			want: &TwoDimensional{matrix: [][]float64{
				{0, 0},
				{0, 0},
			}},
		},
		{
			name: "",
			args: args{
				first: &TwoDimensional{matrix: [][]float64{
					{5, 8, -4},
					{6, 9, -5},
					{4, 7, -3},
				}},
				second: &TwoDimensional{matrix: [][]float64{
					{2},
					{-3},
					{1},
				}},
			},
			want: &TwoDimensional{matrix: [][]float64{
				{-18},
				{-20},
				{-16},
			}},
		},
		{
			name: "",
			args: args{
				first: &TwoDimensional{matrix: [][]float64{
					{5, 8, -4},
					{6, 9, -5},
					{4, 7, -3},
				}},
				second: &TwoDimensional{matrix: [][]float64{
					{3, 2, 5},
					{4, -1, 3},
					{9, 6, 5},
				}},
			},
			want: &TwoDimensional{matrix: [][]float64{
				{11, -22, 29},
				{9, -27, 32},
				{13, -17, 26},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Multiplication(tt.args.first, tt.args.second)
			require.NoError(t, err)
			assertMatrix(t, *tt.want, *got)
		})
	}
}
