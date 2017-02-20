package IntelligenceRepository

import ("math"
	FR "FunctionalRepository"
)

type regFuncs interface {
	RegressionUnivarCoeffs(x []float64, y []float64)(float64, float64)
	RegressionTrivarCoeffs(x1 []float64,x2 []float64,x3 []float64, y []float64)(float64, float64, float64, float64)
}

type regressionFuncs struct{}

func RegressionFuncs() regFuncs{
	return &regressionFuncs{}
}

func (r regressionFuncs)RegressionUnivarCoeffs(x []float64, y []float64)(float64, float64){
	theta0 := 1.0
	theta1 := 0.0
	var theta0I, theta1I, t0, t1 float64
	var dist, minDist, incCnt, min0, min1 float64
	cost := CostFunctions()
	support := FR.SupportFuncs()
	t0 = theta0
	t1 = theta1
	min0 = theta0I
	min1 = theta1I
	for (theta0I != t0 || theta1I != t1){
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
		theta0I, theta1I = cost.DiffUnivarCostFunc(theta0, theta1, x, y)
		t0 = theta0
		t1 = theta1
		theta0 = theta0I
		theta1 = theta1I
	}
	return min0, min1
}

func (r regressionFuncs)RegressionTrivarCoeffs(x1 []float64,x2 []float64,x3 []float64, y []float64)(float64, float64, float64, float64){
	theta0 := 1.0
	theta1 := 0.0
	theta2 := 0.0
	theta3 := 0.0
	var theta0I, theta1I, theta2I, theta3I, t0, t1, t2, t3 float64
	var dist, minDist, incCnt, min0, min1, min2, min3 float64
	cost := CostFunctions()
	support := FR.SupportFuncs()
	t0 = theta0
	t1 = theta1
	t2 = theta2
	t3 = theta3
	min0 = theta0I
	min1 = theta1I
	min2 = theta2I
	min3 = theta3I
	for (theta0I != t0 || theta1I != t1 || theta2I != t2 || theta3I != t3){
		dist = math.Sqrt((t0-theta0I)*(t0-theta0I)+(t1-theta1I)*(t1-theta1I)+(t2-theta2I)*(t2-theta2I)+(t3-theta3I)*(t3-theta3I))
		if(incCnt == 0){
			min0 = theta0
			min1 = theta1
			min2 = theta2
			min3 = theta3
			minDist =dist
		}
		minDistRounded := support.PrecisionRounding(minDist,2)
		if(math.Abs(dist) <math.Abs(minDistRounded) || minDistRounded == 0){
			min0 = theta0I
			min1 = theta1I
			min2 = theta2I
			min3 = theta3I
			minDist = dist
		}else{
			incCnt++
		}
		if(incCnt >5 ||minDistRounded==0){
			break
		}
		theta0I, theta1I, theta2I, theta3I  = cost.DiffTrivarCostFunc(theta0, theta1, theta2, theta3, x1, x2, x3, y)
		t0 = theta0
		t1 = theta1
		t2 = theta2
		t3 = theta3
		theta0 = theta0I
		theta1 = theta1I
		theta2 = theta2I
		theta3 = theta3I
	}
	return min0, min1, min2, min3
}