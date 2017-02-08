package FunctionalRepository

import (
	CR "ConfigurationRepository"
	"strconv"
	"reflect"
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

func GetWeightedFieldArray(valsObj CR.ControlledVals, field string) []float64{
	fldValue := make([] float64,0)
	for _,j := range valsObj.CtrledVals{
		fldValue = append(fldValue,reflect.ValueOf(j).FieldByName(field).Float())//*float64(j.HomesCount))
	}
	return fldValue
}


