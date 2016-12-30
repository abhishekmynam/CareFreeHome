package LogicRepository

import "sort"

func GetOutliers( dataList []float64)[]float64{
	listOutliers:= []float64{}
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

	for i := range dataList{
		if(dataList[len(dataList)-i-1]>= maxForOL){
			listOutliers = append(listOutliers,dataList[len(dataList)-i-1])
		}else{break}
	}

	return listOutliers;
}

