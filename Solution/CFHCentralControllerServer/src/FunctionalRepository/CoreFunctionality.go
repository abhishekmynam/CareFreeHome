package FunctionalRepository

import (CR"ConfigurationRepository"
)

func ProcessDataForLocalCondVals ( currCond CR.CurrCond)(tempMode float64, lightMode float64, musicMode float64){
	var thisCondVals CR.ControlledVals
	var lightForTemp = make([]float64,0)
	var musicForLight = make([]float64,0)
	var musicMode, lightMode float64


	thisCondVals = GetDataCurCond(currCond)
	tempArr:= RemoveOutliers(GetWeightedFieldArray(thisCondVals, "TempIn"))
	lightArr := RemoveOutliers(GetWeightedFieldArray(thisCondVals, "LightIn"))
	musicArr := RemoveOutliers(GetWeightedFieldArray(thisCondVals, "MusicIn"))

	tempMode, _ = GetMode(tempArr)
	lightModeOA, lightModeCnt := GetMode(lightArr)
	musicModeOA, musicModeCnt := GetMode(musicArr)

	for _,j:= range thisCondVals.CtrledVals{
		if(j.TempIn == tempMode[0]){
			lightForTemp = append(lightForTemp,j.LightIn * j.HomesCount)
		}
	}
	lightForTempMode, tempLightModeCnt := GetMode(RemoveOutliers(lightForTemp))

	lightMode = (lightForTempMode[0]*tempLightModeCnt + lightModeOA*lightModeCnt)/(tempLightModeCnt+lightModeCnt)

	for _,j := range thisCondVals.CtrledVals{
		if(j.TempIn == tempMode[0] && j.LightIn == lightMode){
			musicForLight = append(musicForLight, j.MusicIn*j.HomesCount)
		}
	}
	musicForLightMode, lightMusicModeCnt := GetMode(RemoveOutliers(musicForLight))

	musicMode = (musicForLightMode[0]*lightMusicModeCnt + musicModeOA*musicModeCnt)/(musicModeCnt+lightMusicModeCnt)

	return tempMode, lightMode, musicMode
}