package FunctionalRepository

import (
	CR "ConfigurationRepository"
	"strconv"
)

func GetUnProGenDataColObj(tempInside float64, lightInside float64,
 			   musicInside float64,counthomes int64, pplCount int64,
			   outLight float64, outTemp float64, zipCode string, country string, dtime string,
			   working bool, conditionOut string)CR.GlobalCtrlData{

	var gloCtrlData CR.GlobalCtrlData
	var outsidevals CR.OutsideVals
	var insidevals CR.InsideVals

	gloCtrlData.Loc.Zipcode = zipCode
	gloCtrlData.Loc.Country = country

	gloCtrlData.Conddate.Dtime = dtime
	gloCtrlData.Conddate.Working = working

	outsidevals.CondOut = conditionOut
	outsidevals.TempOut = outTemp
	outsidevals.LightOut = outLight
	outsidevals.PplIn = pplCount
	outvals:=make([] CR.OutsideVals,0)
	outvals = append(outvals, outsidevals)

	gloCtrlData.CtrlgVals.CtrlingVals = outvals

	insidevals.TempIn = tempInside
	insidevals.LightIn = lightInside
	insidevals.MusicIn = musicInside
	insidevals.HomesCount = counthomes
	invals := make([] CR.InsideVals,0)
	invals = append(invals, insidevals)

	gloCtrlData.CtrldVals.CtrledVals = invals

	return gloCtrlData

}

func IDGen (prevId string) string{
	var nextId string
	thisId := prevId[3:]
	thisIdNum,_:= strconv.Atoi(thisId)
	thisIdNum  = thisIdNum+1
	thisIdstr := strconv.Itoa(thisIdNum)
	nextId = prevId[:3]+thisIdstr
	return nextId
}

/*var ctrlVals CR.ControllingVals
	var ctrlPpl CR.ControllingPpl
	var ctrlMusic CR.ControllingMusic
	var ctrlLights CR.ControllingLights
	var ctrlTemp CR.ControllingTemp
	var ctrlZip CR.AreaZip
	var ctrlRecTime CR.RecordTime
	var ctrlRecDate CR.RecordDateOfMonth
	var ctrlRecMonth CR.RecordMonth
	var thisControlElem CR.UnProGenDataCol

	ctrlVals.TempInside = tempInside
	ctrlVals.LightsInside = lightInside
	ctrlVals.MusicInside = musicInside
	ctrlVals.CountOfhomes = counthomes
	ctrlvals :=make([] CR.ControllingVals,0)
	ctrlvals = append(ctrlvals, ctrlVals )

	ctrlPpl.PplCount = pplCount
	ctrlPpl.InsideControlVal = ctrlvals
	ctrlppl := make([] CR.ControllingPpl,0)
	ctrlppl = append(ctrlppl, ctrlPpl)

	ctrlMusic.MusicLevel = musicLevel
	ctrlMusic.InsideControlPpl = ctrlppl
	ctrlmusic := make([] CR.ControllingMusic, 0)
	ctrlmusic = append(ctrlmusic, ctrlMusic)

	ctrlLights.OutsideLight = outLight
	ctrlLights.InsideControlMusic = ctrlmusic
	ctrllights := make([] CR.ControllingLights,0)
	ctrllights = append(ctrllights, ctrlLights)

	ctrlTemp.OutsideTemp = outTemp
	ctrlTemp.OutsideControlLight = ctrllights
	ctrltemp := make([] CR.ControllingTemp, 0)
	ctrltemp = append(ctrltemp, ctrlTemp)


	ctrlZip.ZipCode = zipCode
	ctrlZip.OutsideControlTemp = ctrltemp
	ctrlzip := make([] CR.AreaZip, 0)
	ctrlzip = append(ctrlzip, ctrlZip)

	ctrlRecTime.TimeRecord = recTime
	ctrlRecTime.AreaZipToRecord = ctrlzip
	ctrlrectime := make([] CR.RecordTime,0)
	ctrlrectime = append(ctrlrectime, ctrlRecTime)

	ctrlRecDate.DateRecord = recDate
	ctrlRecDate.ControlTime = ctrlrectime
	ctrlrecdate := make([] CR.RecordDateOfMonth,0)
	ctrlrecdate = append(ctrlrecdate, ctrlRecDate)

	ctrlRecMonth.MonthRecord = recMonth
	ctrlRecMonth.ControlDate =ctrlrecdate
	ctrlrecmonth := make([] CR.RecordMonth, 0)
	ctrlrecmonth = append(ctrlrecmonth, ctrlRecMonth)

	thisControlElem.CondOutRecord = conditionOut
	thisControlElem.ControlMonth = ctrlrecmonth


	return  thisControlElem*/