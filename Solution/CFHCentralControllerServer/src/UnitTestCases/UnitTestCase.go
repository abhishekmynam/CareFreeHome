package main

import (
	"FunctionalRepository"
	"fmt"
	"time"
	"ConfigurationRepository"
	"IntelligenceRepository"
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
	//UpdateProcessedCtrldData()
	//ProcessDataForLocalCondVals()
	predict()

}

func getOutliers(){
	outliersList := []float64{3, 5, 3, 11, 6, 7, 4, 6, 9, 0, 4, 5, 21, 22, 23, -12, -13}
	//outliersList := FunctionalRepository.GetOutliers(DataList)
	fmt.Println(outliersList)
}

func RemoveOutliers(){
	FR := FunctionalRepository.Stats()
	outliersList := []float64{3, 5, 3, 11, 6, 7, 4, 6, 9, 0, 4, 5, 21, 22, 23, -12, -13}
	NooutliersList := FR.RemoveOutliers(outliersList)
	fmt.Println(NooutliersList)
}

func getMode(){
	FR := FunctionalRepository.Stats()
	DataList := []float64{3, 5, 3, 11, 6, 7, 4, 6, 9, 0, 4, 5, 21, 22, 23, -12, -13}
	mode,_ := FR.GetMode(DataList)
	fmt.Println(mode)
}

func getStdDev(){
	FR := FunctionalRepository.Stats()
	DataList := []float64{3, 5, 3, 11, 6, 7, 4, 6, 9, 0, 4, 5, 21, 22, 23, -12, -13, 12, 120}
	StdDev := FR.GetStdDev(DataList)
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
	save := FunctionalRepository.InsertIntoDB()
	userSaved:=save.SaveUser(thisUser)
	fmt.Println(userSaved)
}

func removeUser(){
	remove := FunctionalRepository.RemoveFromDb()
	userRemove:= remove.RemoveUser("abhishek.mynam@gmail.com")
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
	//activeUser:= FunctionalRepository.UpdateUser(thisUser)
	//fmt.Println(activeUser)
}

func insertGenData(){
	/*tempInside float64 , lightInside float64,
		musicInside float64, pplCount int64, musicLevel float64,
		outLight float64, outTemp float64, zipCode int64, recTime float64,
		recDate int64, recMonth string, conditionOut string */
	//thisObj:= FunctionalRepository.GetUnProGenDataColObj(15,20,30,50,100,40,55,"19333","USA","10Jan201612",false,"snow")
	//thisStat := FunctionalRepository.UpdateGenControlData(thisObj)
	//thisStatus := FunctionalRepository.SaveNewOutCondition(thisObj)
	//thisStatus := FunctionalRepository.UpdateGenControlData(thisObj)
	//fmt.Println(thisStat)
}

func getIdGen(){
	//id := "LOC1234"
	//thisid:=FunctionalRepository.IDGen(id)
	//fmt.Println(thisid)
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
	//thisArr:=FunctionalRepository.GetWeightedFieldArray(thisVals, "TempIn")
	//fmt.Println(thisArr)

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

	//FunctionalRepository.UpdateProcessedCtrldData(pcd)
}


func ProcessDataForLocalCondVals(){
	var conditions  ConfigurationRepository.CurrCond
	conditions.ZipCode = "19333"
	conditions.Country = "USA"
	conditions.Working = false
	conditions.DateTime = "10Jan201612"
	conditions.Outvals.TempOut = 55
	conditions.Outvals.LightOut = 40
	conditions.Outvals.PplIn = 100
	conditions.Outvals.CondOut = "snow"

	x,y,z:=FunctionalRepository.ProcessDataForLocalCondVals(conditions)
	fmt.Println(x,y,z)
}



type Fooer interface {
	Foo1(s int)int
	Foo2()
	Foo3()
}

type Foo struct {
}

func (f Foo) Foo1(s int) int{
	fmt.Println("Foo1() here",s)
	return s
}

func (f Foo) Foo2() {
	fmt.Println("Foo2() here")
}

func (f Foo) Foo3() {
	fmt.Println("Foo3() here")
}

func NewFoo() Fooer {
	return &Foo{}
}

func predict(){
	thisX := 60.0
	thisY:=0.0
	x:= []float64{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,24,25,26,27,28,29,30,31,32,33,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50}
	y:=[]float64{5,7,8,6,8,9,10,12,15,14,15,16,17,18,19,17,20,21,22,22,25,26,30,31,31,31,35,32,36,33,34,37,40,40,41,43,42,46,46,43,49,48,48,50,53,54,56,60}
	reg := IntelligenceRepository.RegressionFuncs()
	t0, t1 := reg.RegressionLinearCoeffs(x,y)
	hyp := IntelligenceRepository.HypothesisGen()
	thisY = hyp.UnivarRegHypothesis(t0, t1, thisX)
	fmt.Println(thisY)


}

func testInterface(){
	f := NewFoo()

	t := f.Foo1(77)
	fmt.Println(t)

	f.Foo2()

	f.Foo3()
}
