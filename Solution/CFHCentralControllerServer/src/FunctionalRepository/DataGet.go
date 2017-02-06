package FunctionalRepository

import(
	CR "ConfigurationRepository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetDataCurCond (CurrCond CR.CurrCond) CR.ControlledVals{
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		panic(err)
	}
	defer session.Close()

	locColl := session.DB(CR.DBInstance).C(CR.LocColl)
	dtimeColl := session.DB(CR.DBInstance).C(CR.CondDateColl)
	ctrlgValsColl := session.DB(CR.DBInstance).C(CR.ControllingValsColl)
	ctrldValsColl := session.DB(CR.DBInstance).C(CR.ControlledValsColl)

	var locObj CR.Location
	var dtimeObj CR.CondDate
	var ctrlngValsObj CR.ControllingVals
	var controllingId string
	var ctrldValsObj CR.ControlledVals

	locColl.Find(bson.M{"zipcode":CurrCond.ZipCode, "country":CurrCond.Country}).Select(bson.M{"zipcode":1,"locid":1}).One(&locObj)
	dtimeColl.Find(bson.M{"dtime":CurrCond.DateTime, "working":CurrCond.Working, "locid":locObj.LocId}).Select(
		bson.M{"dtime":1, "conddateid":1}).One(&dtimeObj )
	ctrlgValsColl.Find(bson.M{"conddateid":dtimeObj.CondDateId,
		"ctrlingvals.condout":CurrCond.Outvals.CondOut,
		"ctrlingvals.tempout":CurrCond.Outvals.TempOut,
		"ctrlingvals.lightout":CurrCond.Outvals.LightOut,
		"ctrlingvals.pplin":CurrCond.Outvals.PplIn}).Select(
		bson.M{"ctrlingvals":1}).One(&ctrlngValsObj)
	for _,j:= range ctrlngValsObj.CtrlingVals{
		if(j.CondOut == CurrCond.Outvals.CondOut &&
			j.TempOut == CurrCond.Outvals.TempOut &&
			j.LightOut == CurrCond.Outvals.LightOut &&
			j.PplIn == CurrCond.Outvals.PplIn){
			controllingId = j.CtrlValsId
			break
		}
	}
	ctrldValsColl.Find(bson.M{"ctrlvalsid":controllingId}).Select(bson.M{"ctrledvals":1}).One(&ctrldValsObj)

	return ctrldValsObj

}