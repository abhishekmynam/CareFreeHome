package FunctionalRepository

import (
	"sort"
	"math"
//	"fmt"
)
type stats interface {
	GetMode(numbers []float64)([]float64,int64)
	GetStdDev(numbers []float64)float64
	GetVariance(numbers []float64)float64
	GetMean(numbers []float64)float64
	SumOfSliceElements(numbers []float64)float64
	RemoveOutliers(numbers []float64)[]float64
}
type statFuncs struct {
}
func Stats() stats{
	return &statFuncs{}
}

func getOutliers( dataList []float64)(float64 ,float64){
	var minOutlier  float64 =-999
	var maxOutlier  float64 =-999
	listOutliers:= []float64{}
	listOutliers2:= []float64{}
	sort.Float64s(dataList)

	var dataLength,dataLenBy2 int
	var Q1,Q3,minForOL,maxForOL float64

	dataLength = len(dataList)

	if(dataLength%2==0){
		dataLenBy2= dataLength/2

		if(dataLenBy2%2==0) {
			Q1 = (dataList[dataLenBy2 / 2] + dataList[dataLenBy2 + 1]) / 2
			Q3 = (dataList[dataLenBy2 + dataLenBy2 / 2] + dataList[dataLenBy2/2 + dataLenBy2 + 1]) / 2
		}else{
			Q1 = dataList[1/2+dataLenBy2/2]
			Q3 = dataList[dataLenBy2 + 1/2 + dataLenBy2/2]
		}
	}else{
		dataLenBy2 = 1/2+ dataLength/2

		if(dataLenBy2%2==0) {
			Q1 = (dataList[dataLenBy2 / 2] + dataList[dataLenBy2 + 1]) / 2
			Q3 = (dataList[dataLenBy2 + dataLenBy2 / 2] + dataList[dataLenBy2/2 + dataLenBy2 + 1]) / 2
		}else{
			Q1 = dataList[1/2+dataLenBy2/2]
			Q3 = dataList[dataLenBy2 + 1/2 + dataLenBy2/2]
		}
	}

	minForOL = Q1 - 1.5*(Q3-Q1)
	maxForOL = Q3 + 1.5*(Q3-Q1)

	for i := range dataList{
		if (dataList[i]<=minForOL){
			listOutliers = append(listOutliers,dataList[i])
		}else{break}
	}
	sort.Float64s(listOutliers)
	if(len(listOutliers)!=0){
		minOutlier =  listOutliers[0]
	}
	for i := range dataList{
		if(dataList[len(dataList)-i-1]>= maxForOL){
			listOutliers2 = append(listOutliers2,dataList[len(dataList)-i-1])
		}else{break}
	}
	sort.Float64s(listOutliers2)
	if(len(listOutliers2)!=0){
		maxOutlier =  listOutliers2[0]
	}
	return minOutlier,maxOutlier;
}

func (s statFuncs)GetMode(numbers []float64) (modes []float64, highestFrequency int64) {
	frequencies := make(map[float64]int64, len(numbers))
	//highestFrequency := 0
	for _, x := range numbers {
		frequencies[x]++
		if frequencies[x] > highestFrequency {
			highestFrequency = frequencies[x]
		}
	}
	for x, frequency := range frequencies {
		if frequency == highestFrequency {
			modes = append(modes, x)
		}
	}
	if highestFrequency == 1 || len(modes) == len(numbers) {
		modes = modes[:0] // Or: modes = []float64{}
	}
	sort.Float64s(modes)
	return modes, highestFrequency
}

func (s statFuncs)GetStdDev(numbers []float64) float64 {

	variance := s.GetVariance(numbers) / float64(len(numbers)-1)
	return math.Sqrt(variance)
}

func (s statFuncs)GetVariance(numbers []float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += math.Pow(number-s.GetMean(numbers), 2)
	}
	variance := total / float64(len(numbers)-1)
	return variance
}

func (s statFuncs)GetMean(numbers []float64)float64{
	mean := s.SumOfSliceElements(numbers) / float64(len(numbers))
	return mean
}

func (s statFuncs)SumOfSliceElements(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return total
}

func(s statFuncs) RemoveOutliers (numbers []float64)[]float64{
	//var outliersRemoved = make([] float64,0)
	minOutliers, maxOutliers := getOutliers(numbers)
	sort.Float64s(numbers)
	start:= 0
	end:= len(numbers)
	//numbers[sort.SearchFloat64s(numbers, minOutliers)+1:sort.SearchFloat64s(numbers, maxOutliers)]
	if(minOutliers != -999){
		start= sort.SearchFloat64s(numbers, minOutliers)+1
	}
	if(maxOutliers != -999) {
		end = sort.SearchFloat64s(numbers, maxOutliers)
	}
	return numbers[start:end]
}
