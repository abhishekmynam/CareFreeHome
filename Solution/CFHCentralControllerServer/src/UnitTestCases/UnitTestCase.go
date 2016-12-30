package main

import (
	"LogicRepository"
	"fmt"

)

func main() {
	DataList := []float64{3, 5, 3, 11, 6, 7, 4, 6, 9, 0, 4, 5, 21, 22, 23, -12, -13}
	outliersList :=LogicRepository.GetOutliers(DataList)
	fmt.Println(outliersList)
}

