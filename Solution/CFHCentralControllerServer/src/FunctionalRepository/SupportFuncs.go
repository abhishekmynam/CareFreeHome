package FunctionalRepository

import "ConfigurationRepository"

func GetUnProGenDataColObj(tempInside float64, lightInside float64,
 			   musicInside float64,counthomes int64, pplCount int64, musicLevel float64,
			   outLight float64, outTemp float64, zipCode int64, recTime float64,
			   recDate int64, recMonth string, conditionOut string)ConfigurationRepository.UnProGenDataCol{
	var ctrlVals ConfigurationRepository.ControllingVals
	var ctrlPpl ConfigurationRepository.ControllingPpl
	var ctrlMusic ConfigurationRepository.ControllingMusic
	var ctrlLights ConfigurationRepository.ControllingLights
	var ctrlTemp ConfigurationRepository.ControllingTemp
	var ctrlZip ConfigurationRepository.AreaZip
	var ctrlRecTime ConfigurationRepository.RecordTime
	var ctrlRecDate ConfigurationRepository.RecordDateOfMonth
	var ctrlRecMonth ConfigurationRepository.RecordMonth
	var thisControlElem ConfigurationRepository.UnProGenDataCol

	ctrlVals.TempInside = tempInside
	ctrlVals.LightsInside = lightInside
	ctrlVals.MusicInside = musicInside
	ctrlVals.CountOfhomes = counthomes
	//ctrlvals :=make([] ConfigurationRepository.ControllingVals,0)
	//ctrlvals = append(ctrlvals, ctrlVals )

	ctrlPpl.PplCount = pplCount
	ctrlPpl.InsideControlVal = ctrlVals
	//ctrlppl := make([] ConfigurationRepository.ControllingPpl,0)
	//ctrlppl = append(ctrlppl, ctrlppl)

	ctrlMusic.MusicLevel = musicLevel
	ctrlMusic.InsideControlPpl = ctrlPpl

	ctrlLights.OutsideLight = outLight
	ctrlLights.InsideControlMusic = ctrlMusic

	ctrlTemp.OutsideTemp = outTemp
	ctrlTemp.OutsideControlLight = ctrlLights

	ctrlZip.ZipCode = zipCode
	ctrlZip.OutsideControlTemp = ctrlTemp

	ctrlRecTime.TimeRecord = recTime
	ctrlRecTime.AreaZipToRecord = ctrlZip

	ctrlRecDate.DateRecord = recDate
	ctrlRecDate.ControlTime = ctrlRecTime

	ctrlRecMonth.MonthRecord = recMonth
	ctrlRecMonth.ControlDate =ctrlRecDate

	thisControlElem.CondOutRecord = conditionOut
	thisControlElem.ControlMonth = ctrlRecMonth

	return  thisControlElem
}