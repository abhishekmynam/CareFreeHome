package IntelligenceRepository

type hypothesisMethods interface {
	UnivarRegHypothesis(theta0 float64, theta1 float64, x float64 )float64
	TrivarRegHypothesis(theta0 float64, theta1 float64, theta2 float64, theta3 float64, x1 float64, x2 float64, x3 float64 )float64
}

type hypothesis struct{
}

func HypothesisGen() hypothesisMethods{
	return &hypothesis{}
}

func (h hypothesis) UnivarRegHypothesis(theta0 float64, theta1 float64, x float64 )float64{
	var hThetaX float64
	hThetaX = theta0 + theta1*x
	return hThetaX
}

func (h hypothesis) TrivarRegHypothesis(theta0 float64, theta1 float64, theta2 float64, theta3 float64, x1 float64, x2 float64, x3 float64 )float64{
	var hThetaX float64
	hThetaX = theta0 + theta1*x1 + theta2*x2 + theta3*x3
	return hThetaX
}