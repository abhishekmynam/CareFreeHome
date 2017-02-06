package main

import (
	"FunctionalRepository"
	"fmt"
	"time"
	"ConfigurationRepository"
)

func main() {
	//getOutliers()
	//errorLogFileCheck()
	//getMode()
	//getStdDev()
	//saveUser()
	//removeUser()
	//updateUser()
	//insertGenData()
	//getIdGen()
	//removeOutliers()
	//getDataForCondition()
	UpdateProcessedCtrldData()

}

func getOutliers(){
	outliersList := []float64{3, 5, 3, 11, 6, 7, 4, 6, 9, 0, 4, 5, 21, 22, 23, -12, -13}
	//outliersList := FunctionalRepository.GetOutliers(DataList)
	fmt.Println(outliersList)
}

func removeOutliers(){
	outliersList := []float64{3, 5, 3, 11, 6, 7, 4, 6, 9, 0, 4, 5, 21, 22, 23, -12, -13}
	NooutliersList := FunctionalRepository.RemoveOutliers(outliersList)
	fmt.Println(NooutliersList)
}

func getMode(){
	DataList := []float64{3, 5, 3, 11, 6, 7, 4, 6, 9, 0, 4, 5, 21, 22, 23, -12, -13}
	mode,_ := FunctionalRepository.GetMode(DataList)
	fmt.Println(mode)
}

func getStdDev(){
	DataList := []float64{3, 5, 3, 11, 6, 7, 4, 6, 9, 0, 4, 5, 21, 22, 23, -12, -13, 12, 120}
	StdDev := FunctionalRepository.GetStdDev(DataList)
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
	thisUser.Email="abhishek15.mynam@gmail.com"
	thisUser.PwdConfirm = "thisispassword1"
	thisUser.Ph_no =2157918506
	thisUser.Address="beech street1"
	thisUser.City="Maywood1"
	thisUser.Zipcode=07646
	thisUser.State = "NJ"
	userSaved:=FunctionalRepository.SaveUser(thisUser)
	fmt.Println(userSaved)
}

func removeUser(){
	userRemove:= FunctionalRepository.RemoveUser("abhishek.mynam@gmail.com")
	fmt.Println(userRemove)
}

func updateUser(){
	var thisUser ConfigurationRepository.User
	thisUser.UsrFName ="Abhishek1"
	thisUser.UsrLName="Mynam1"
	thisUser.Email="abhishek13.mynam@gmail.com"
	thisUser.PwdConfirm = "thisispassword1"
	thisUser.Ph_no =2157918506
	thisUser.Address="beech street1"
	thisUser.City="Maywood1"
	thisUser.Zipcode=07646
	thisUser.State = "NJ"
	activeUser:= FunctionalRepository.UpdateUser(thisUser)
	fmt.Println(activeUser)
}

func insertGenData(){
	/*tempInside float64 , lightInside float64,
		musicInside float64, pplCount int64, musicLevel float64,
		outLight float64, outTemp float64, zipCode int64, recTime float64,
		recDate int64, recMonth string, conditionOut string */
	thisObj:= FunctionalRepository.GetUnProGenDataColObj(15,20,30,50,100,40,55,"19333","USA","10Jan201612",false,"snow")
	thisStat := FunctionalRepository.UpdateGenControlData(thisObj)
	//thisStatus := FunctionalRepository.SaveNewOutCondition(thisObj)
	//thisStatus := FunctionalRepository.UpdateGenControlData(thisObj)
	fmt.Println(thisStat)
}

func getIdGen(){
	id := "LOC1234"
	thisid:=FunctionalRepository.IDGen(id)
	fmt.Println(thisid)
}
func getDataForCondition(){
	var conditions  ConfigurationRepository.CurrCond
	var thisVals ConfigurationRepository.ControlledVals
	conditions.ZipCode = "19333"
	conditions.Country = "USA"
	conditions.Working = false
	conditions.DateTime = "10Jan201612"
	conditions.Outvals.TempOut = 55
	conditions.Outvals.LightOut = 40
	conditions.Outvals.PplIn = 100
	conditions.Outvals.CondOut = "snow"

	thisVals = FunctionalRepository.GetDataCurCond(conditions)
	getAsArray(thisVals)
}

func getAsArray (thisVals ConfigurationRepository.ControlledVals){
	thisArr:=FunctionalRepository.GetWeightedFieldArray(thisVals, "TempIn")
	fmt.Println(thisArr)

}

func UpdateProcessedCtrldData (){
	var pcd ConfigurationRepository.ProcessedCtrlData
	pcd.Zipcode ="19333"
	pcd.Country = "USA"
	pcd.Dtime = "12Jan201712"
	pcd.Working = false
	pcd.CondOut = "snow"
	pcd.TempOut = 55
	pcd.LightOut = 65
	pcd.PplIn = 100
	pcd.TempIn = 46
	pcd.LightIn = 36
	pcd.MusicIn = 16

	FunctionalRepository.UpdateProcessedCtrldData(pcd)
}