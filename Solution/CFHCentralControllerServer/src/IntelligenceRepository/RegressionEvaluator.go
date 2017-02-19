package IntelligenceRepository

import ("math"
	FR "FunctionalRepository"
)

type regFuncs interface {
	RegressionLinearCoeffs(x []float64, y []float64)(float64, float64)
}

type regressionFuncs struct{}

func RegressionFuncs() regFuncs{
	return &regressionFuncs{}
}

func (r regressionFuncs)RegressionLinearCoeffs(x []float64, y []float64)(float64, float64){
	theta0 := 1.0
	theta1 := 0.0
	var theta0I, theta1I, t0, t1 float64
	var dist, minDist, incCnt, min0, min1 float64
	cost := CostFunctions()
	support := FR.SupportFuncs()
	t0 = theta0
	t1= theta1
	min0 = theta0I
	min1 = theta1I
	for (theta0I!= t0 || theta1I !=t1){
		dist = math.Sqrt((t0-theta0I)*(t0-theta0I)+(t1-theta1I)*(t1-theta1I))
		if(incCnt == 0){minDist =dist}
		minDistRounded := support.PrecisionRounding(minDist,2)
		if(math.Abs(dist) <math.Abs(minDistRounded) || minDistRounded == 0){
			min0 = theta0I
			min1 = theta1I
			minDist = dist
		}else{
			incCnt++
		}
		if(incCnt >5 ||minDistRounded==0){
			break
		}
		theta0I, theta1I = cost.DiffLinearCostFunc(theta0, theta1, x, y)
		t0 = theta0
		t1= theta1
		theta0 = theta0I
		theta1 = theta1I
	}
	return min0, min1
}