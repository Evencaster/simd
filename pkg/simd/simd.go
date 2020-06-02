package simd

import "github.com/illfate2/simd/pkg/matrix"

type Processor struct {
	ProcessorArgs
}

type ProcessorArgs struct {
	A *matrix.TwoDimensional
	B *matrix.TwoDimensional
	E *matrix.TwoDimensional
	G *matrix.TwoDimensional
	K int
}

func NewProcessor(args ProcessorArgs) Processor {
	return Processor{args}

}

func (p Processor) Result() (*matrix.TwoDimensional, error) {
	firstRes, err := p.firstSumOfResult()
	if err != nil {
		return nil, err
	}
	secondRes, err := p.secondSumOfResult()
	if err != nil {
		return nil, err
	}
	err = firstRes.Sum(secondRes)
	return firstRes, err
}

func (p Processor) fThreeDimensional() (*matrix.TwoDimensional, error) {
	res, err := p.firstSumOfThreeDem()
	if err != nil {
		return nil, err
	}
	secondRes, err := p.secondSumOfThreeDem()
	if err != nil {
		return nil, err
	}
	err = res.Sum(secondRes)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p Processor) dThreeDimensional() (*matrix.TwoDimensional, error) {
	return matrix.Multiplication(p.A, p.B)
}

func (p Processor) evaluateKTimes(dimensional *matrix.TwoDimensional) (*matrix.TwoDimensional, error) {
	var err error
	dupl := dimensional.Copy()
	for i := 0; i < p.K; i++ {
		dupl, err = matrix.Multiplication(dupl, dupl)
		if err != nil {
			return nil, err
		}
	}
	return dupl, nil
}

func (p Processor) firstSumOfResult() (*matrix.TwoDimensional, error) {
	res, err := p.fijk()
	if err != nil {
		return nil, err
	}

	duplG := p.G.Copy()
	duplG.MultiplicationInt(3)

	err = duplG.Sum(matrix.NewTwoDimensionalWithValue(duplG.M(), duplG.N(), -2))
	if err != nil {
		return nil, err
	}

	resMult, err := matrix.Multiplication(res, duplG)
	if err != nil {
		return nil, err
	}
	resMult, err = matrix.Multiplication(resMult, p.G.Copy())
	return resMult, err
}

func (p Processor) secondSumOfResult() (*matrix.TwoDimensional, error) {
	vDijk, err := p.vDijk()
	if err != nil {
		return nil, err
	}

	fijk, err := p.fijk()
	if err != nil {
		return nil, err
	}
	mult := matrix.MustMultiplication(fijk, vDijk)
	mult.MultiplicationInt(4)
	duplDijk := vDijk.Copy()
	duplDijk.MultiplicationInt(3)
	err = mult.Sum(matrix.Negative(duplDijk))
	if err != nil {
		return nil, err
	}
	mult = matrix.MustMultiplication(mult, p.G)

	sum := matrix.MustSum(vDijk, mult)
	gSum := matrix.MustSum(matrix.NewTwoDimensionalWithValue(p.G.M(), p.G.N(), 1), p.G)
	return matrix.MustMultiplication(sum, gSum), nil
}

func (p Processor) firstSumOfThreeDem() (*matrix.TwoDimensional, error) {
	multRes, err := matrix.SumInt(p.E, 2)
	if err != nil {
		return nil, err
	}
	sumRes, err := matrix.SumInt(multRes, -1)
	if err != nil {
		return nil, err
	}
	multRes = matrix.MultiplicationInt(sumRes, convRes)
	multRes, err = matrix.Multiplication(multRes, p.E)
	if err != nil {
		return nil, err
	}
	return multRes, nil
}

func (p Processor) secondSumOfThreeDem() (*matrix.TwoDimensional, error) {

	mult := matrix.MultiplicationInt(p.E, multInt)

	one := matrix.NewTwoDimensionalWithValue(mult.M(), mult.N(), 1)
	err = one.Sum(mult)
	if err != nil {
		return nil, err
	}

	multRes := matrix.MultiplicationInt(one, bToA)

	oneE := matrix.NewTwoDimensionalWithValue(p.E.M(), p.E.N(), 1)
	err = oneE.Sum(matrix.Negative(p.E))
	if err != nil {
		return nil, err
	}
	return matrix.Multiplication(multRes, oneE)
}

func (p Processor) vDijk() (*matrix.TwoDimensional, error) {
	// 1 - P(1-Dijk)
	dijk, err := p.dThreeDimensional()
	if err != nil {
		return nil, err
	}
	one := matrix.NewTwoDimensionalWithValue(dijk.M(), dijk.N(), 1)
	err = one.Sum(matrix.Negative(dijk))
	if err != nil {
		return nil, err
	}
	evaluated, err := p.evaluateKTimes(one)
	if err != nil {
		return nil, err
	}
	newOne := matrix.NewTwoDimensionalWithValue(dijk.M(), dijk.N(), 1)
	err = newOne.Sum(matrix.Negative(evaluated))
	return newOne, err
}

func (p Processor) fijk() (*matrix.TwoDimensional, error) {
	fThreeDimensional, err := p.fThreeDimensional()
	if err != nil {
		return nil, err
	}
	res, err := p.evaluateKTimes(fThreeDimensional)
	return res, err
}
