package IntelligenceRepository

type costFuncs interface {
	LinearCostFunction(theta0 float64, theta1 float64, x []float64, y []float64) float64
	DiffLinearCostFunc(theta0 float64, theta1 float64, x []float64, y []float64)(float64, float64)
}
type costFunction struct{}

func CostFunctions() costFuncs{
	return &costFunction{}
}

func (c costFunction)LinearCostFunction(theta0 float64, theta1 float64, x []float64, y []float64) float64{
	var cost,sqError float64

	hyp := HypothesisGen()

	for i,j := range x{
		h := hyp.UnivarRegHypothesis(theta0, theta1,j)
		thisP := (h-y[i])*(h-y[i])
		sqError = sqError + thisP
	}
	cost = sqError/(2*float64(len(x)))

	return cost
}
func (c costFunction)DiffLinearCostFunc(theta0 float64, theta1 float64, x []float64, y []float64)(float64, float64){
	var sum0, sum1, theta0R, theta1R float64
	alpha := 0.001
	hyp := HypothesisGen()
	for i, j := range x{
		h:= hyp.UnivarRegHypothesis(theta0, theta1, j)
		t0 := h - y[i]
		t1 := (h-y[i])*j
		sum0 = sum0+t0
		sum1 = sum1+t1
	}
	theta0R = theta0 - alpha*(sum0/float64(len(x)))
	theta1R = theta1 - alpha*(sum1/float64(len(x)))

	return theta0R, theta1R
}