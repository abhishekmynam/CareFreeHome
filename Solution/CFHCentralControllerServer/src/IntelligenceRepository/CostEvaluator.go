package IntelligenceRepository

type costFuncs interface {
	LinearCostFunction(theta0 float64, theta1 float64, x []float64, y []float64) float64
	DiffUnivarCostFunc(theta0 float64, theta1 float64, x []float64, y []float64)(float64, float64)
	DiffTrivarCostFunc(theta0 float64, theta1 float64, theta2 float64, theta3 float64,
			   x1 []float64, x2 []float64, x3 []float64, y []float64)(float64, float64, float64, float64)
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
func (c costFunction)DiffUnivarCostFunc(theta0 float64, theta1 float64, x []float64, y []float64)(float64, float64){
	var sum0, sum1, theta0R, theta1R float64
	alpha := 0.001
	hyp := HypothesisGen()
	for i, _ := range x{
		h:= hyp.UnivarRegHypothesis(theta0, theta1, x[i])
		t0 := h - y[i]
		t1 := (h-y[i])*x[i]
		sum0 = sum0+t0
		sum1 = sum1+t1
	}
	theta0R = theta0 - alpha*(sum0/float64(len(x)))
	theta1R = theta1 - alpha*(sum1/float64(len(x)))

	return theta0R, theta1R
}

func (c costFunction)DiffTrivarCostFunc(theta0 float64, theta1 float64, theta2 float64, theta3 float64,
					x1 []float64, x2 []float64, x3 []float64, y []float64)(float64, float64, float64, float64){
	var sum0, sum1, sum2, sum3, theta0R, theta1R, theta2R, theta3R float64
	alpha := 0.0003
	hyp := HypothesisGen()
	for i, _ := range x1{
		h:= hyp.TrivarRegHypothesis(theta0, theta1, theta2, theta3, x1[i], x2[i], x3[i] )
		t0 := h - y[i]
		t1 := (h-y[i])*x1[i]
		t2 := (h-y[i])*x2[i]
		t3 := (h-y[i])*x3[i]
		sum0 = sum0+t0
		sum1 = sum1+t1
		sum2 = sum2+t2
		sum3 = sum3+t3
	}
	theta0R = theta0 - alpha*(sum0/float64(len(x1)))
	theta1R = theta1 - alpha*(sum1/float64(len(x1)))
	theta2R = theta2 - alpha*(sum2/float64(len(x1)))
	theta3R = theta3 - alpha*(sum3/float64(len(x1)))

	return theta0R, theta1R, theta2R, theta3R
}