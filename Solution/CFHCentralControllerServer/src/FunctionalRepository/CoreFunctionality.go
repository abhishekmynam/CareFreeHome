package FunctionalRepository

import (CR"ConfigurationRepository"
)

func ProcessDataForLocalCondVals ( currCond CR.CurrCond)( float64,float64, float64){
	var thisCondVals CR.ControlledVals
	var lightForTemp = make([]float64,0)
	var musicForLight = make([]float64,0)
	var musicMode, lightMode float64
	var tempLightModeCnt, lightMusicModeCnt int64
	var musicForLightModeVal, lightForTempModeVal float64

	thisCondVals = GetDataCurCond(currCond)
	tempArr:= RemoveOutliers(GetWeightedFieldArray(thisCondVals, "TempIn"))
	lightArr := RemoveOutliers(GetWeightedFieldArray(thisCondVals, "LightIn"))
	musicArr := RemoveOutliers(GetWeightedFieldArray(thisCondVals, "MusicIn"))

	tempMode, _ := GetMode(tempArr)
	lightModeOA, lightModeCnt := GetMode(lightArr)
	musicModeOA, musicModeCnt := GetMode(musicArr)
	controlledValues := thisCondVals.CtrledVals
	for _,j := range controlledValues{
		if(j.TempIn == tempMode[0]){
			lightForTemp = append(lightForTemp,j.LightIn )//* float64(j.HomesCount))
		}
	}
	if(len(lightForTemp)!=0) {
		lightForTempMode, tempLightModeCnt1 := GetMode(lightForTemp)
		lightForTempModeVal = lightForTempMode[0]
		tempLightModeCnt = tempLightModeCnt1
	}
	lightMode = (lightForTempModeVal * float64(tempLightModeCnt) + lightModeOA[0]* float64(lightModeCnt))/float64(tempLightModeCnt+lightModeCnt)

	for _,j := range controlledValues{
		if(j.TempIn == tempMode[0] && j.LightIn == lightMode){
			musicForLight = append(musicForLight, j.MusicIn)//* float64(j.HomesCount))
		}
	}
	if(len(musicForLight)!=0) {
		musicForLightMode, lightMusicModeCnt1 := GetMode(musicForLight)
		musicForLightModeVal = musicForLightMode[0]
		lightMusicModeCnt =lightMusicModeCnt1
	}
	musicMode = (musicForLightModeVal*float64(lightMusicModeCnt )+ musicModeOA[0]*float64(musicModeCnt))/float64(musicModeCnt+lightMusicModeCnt)

	return tempMode[0], lightMode, musicMode
}