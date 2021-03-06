package FunctionalRepository

import(
	CR "ConfigurationRepository"
)

type processData interface{
	ProcessDataForThisHour() string
}

type processDataFuncs struct{
}

func ProcessDBData() processData{
	return &processDataFuncs{}
}

func (p processDataFuncs)ProcessDataForThisHour()string{

	var outsideCondition CR.CurrCond
	var updateData CR.ProcessedCtrlData

	outsideCondition = GetOutsideVals();

	temp, light, music := ProcessDataForLocalCondVals(outsideCondition)

	updateData.Zipcode =outsideCondition.ZipCode
	updateData.Country = outsideCondition.Country
	updateData.Dtime = outsideCondition.DateTime
	updateData.Working = outsideCondition.Working
	updateData.CondOut = outsideCondition.Outvals.CondOut
	updateData.TempOut = outsideCondition.Outvals.TempOut
	updateData.LightOut = outsideCondition.Outvals.LightOut
	updateData.PplIn = outsideCondition.Outvals.PplIn
	updateData.TempIn = temp
	updateData.LightIn = light
	updateData.MusicIn = music
	dbUpdate:= DBDataUpdate()
	procState:= dbUpdate.UpdateProcessedCtrldData(updateData)

	return procState
}
