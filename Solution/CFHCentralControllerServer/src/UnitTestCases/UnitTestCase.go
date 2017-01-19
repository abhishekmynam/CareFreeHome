package main

import (
	"LogicRepository"
	"fmt"
	"time"
	"ConfigurationRepository"
)

func main() {
	//getOutliers()
	//errorLogFileCheck()
	getMode()
}

func getOutliers(){
	DataList := []float64{3, 5, 3, 11, 6, 7, 4, 6, 9, 0, 4, 5, 21, 22, 23, -12, -13}
	outliersList :=LogicRepository.GetOutliers(DataList)
	fmt.Println(outliersList)
}

func getMode(){
	DataList := []float64{3, 5, 3, 11, 6, 7, 4, 6, 9, 0, 4, 5, 21, 22, 23, -12, -13}
	mode :=LogicRepository.GetMode(DataList)
	fmt.Println(mode)
}

func errorLogFileCheck(){
	var message = time.Now().String()+" -- this is check for error logs 2"
	ConfigurationRepository.LogError(message)
}

