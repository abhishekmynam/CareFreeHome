package IntelligenceRepository

type hypothesisMethods interface {
	UnivarRegHypothesis(theta0 float64, theta1 float64, x float64 )float64
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