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
	//getMode()
	//getStdDev()
	saveUser()
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

func getStdDev(){
	DataList := []float64{3, 5, 3, 11, 6, 7, 4, 6, 9, 0, 4, 5, 21, 22, 23, -12, -13, 12, 120}
	StdDev :=LogicRepository.GetStdDev(DataList)
	fmt.Println(StdDev)
}

func errorLogFileCheck(){
	var message = time.Now().String()+" -- this is check for error logs 2"
	ConfigurationRepository.LogError(message)
}

func saveUser(){
	var thisUser ConfigurationRepository.User
	thisUser.UsrFName ="Abhishek1"
	thisUser.UsrLName="Mynam1"
	thisUser.Email="abhishek12.mynam@gmail.com"
	thisUser.PwdConfirm = "thisispassword1"
	thisUser.Ph_no =2157918506
	thisUser.Address="beech street1"
	thisUser.City="Maywood1"
	thisUser.Zipcode=07646
	thisUser.State = "NJ"
	LogicRepository.SaveUser(thisUser)
}

