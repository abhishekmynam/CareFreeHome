package FunctionalRepository

import (
	"sort"
	"math"
	"fmt"
)

func getOutliers( dataList []float64)[]float64{
	minMaxOutlier := []float64{}
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
	minMaxOutlier = append(minMaxOutlier, listOutliers[len(listOutliers)-1] )
	for i := range dataList{
		if(dataList[len(dataList)-i-1]>= maxForOL){
			listOutliers2 = append(listOutliers2,dataList[len(dataList)-i-1])
		}else{break}
	}
	sort.Float64s(listOutliers2)
	minMaxOutlier = append(minMaxOutlier, listOutliers2[0] )
	return minMaxOutlier;
}

func GetMode(numbers []float64) (modes []float64, highestFrequency int64) {
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

func GetStdDev(numbers []float64) float64 {

	variance := GetVariance(numbers) / float64(len(numbers)-1)
	return math.Sqrt(variance)
}

func GetVariance(numbers []float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += math.Pow(number-GetMean(numbers), 2)
	}
	variance := total / float64(len(numbers)-1)
	return variance
}

func GetMean(numbers []float64)float64{
	mean := SumOfSliceElements(numbers) / float64(len(numbers))
	return mean
}

func SumOfSliceElements(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return total
}

func RemoveOutliers (numbers []float64)[]float64{
	var outliersRemoved = make([] float64,0)
	outliers := getOutliers(numbers)
	sort.Float64s(numbers)
	fmt.Println(numbers[sort.SearchFloat64s(numbers, outliers[0])+1:sort.SearchFloat64s(numbers, outliers[1])])
	return outliersRemoved
}
