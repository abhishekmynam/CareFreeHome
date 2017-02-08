package FunctionalRepository

import (
	"gopkg.in/mgo.v2"
	"log"
	CR "ConfigurationRepository"
	"gopkg.in/mgo.v2/bson"
	"time"
	"strconv"
)

func UpdateUser (newUser CR.User) string {
	var DeactivateUserStatus string
	session, err := mgo.Dial(CR.DBserver)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	userColl := session.DB(CR.DBInstance).C(CR.UserMasterCollection)
	var existingUser CR.User
	err = userColl.Find(bson.M{"email":newUser.Email}).One(&existingUser)
	if (len(existingUser.Email)==0){
		DeactivateUserStatus = "User doesn't exist to flip status"
	}else if(len(existingUser.Email)!=0){
		err = userColl.Update(bson.M{"email":newUser.Email},&CR.User{UsrFName: newUser.UsrFName, UsrLName:newUser.UsrLName,
			Email:newUser.Email, PwdConfirm:newUser.PwdConfirm, Ph_no: newUser.Ph_no, Address:newUser.Address,
			City:newUser.City, Zipcode:newUser.Zipcode, State:newUser.State, Country:newUser.Country, UserStatus:newUser.UserStatus, LastUpdated:time.Now()})
		DeactivateUserStatus="User status flipped"
	}
	if err!=nil{
		log.Fatal(err)
		DeactivateUserStatus ="error occured check the logs"
	}
	return DeactivateUserStatus
}

func UpdateProcessedCtrldData(procData CR.ProcessedCtrlData) (updateProcCtrlDataStat string){
	session, err := mgo.Dial(CR.DBserver)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	procDataColl := session.DB(CR.DBInstance).C(CR.ProcessedCtrlDataCollection)
	_,err = procDataColl.Upsert(bson.M{"zipcode":procData.Zipcode, "country":procData.Country,
				"dtime":procData.Dtime, "working":procData.Working, "condout": procData.CondOut,
				"tempout":procData.TempOut, "lightout":procData.LightOut, "pplin":procData.PplIn},procData )
	if err!=nil{
		log.Fatal(err)
	}else{updateProcCtrlDataStat = "success"}

	return updateProcCtrlDataStat
}

func UpdateGenControlData (updateData CR.GlobalCtrlData)string{
	var updateCtrlData string
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
	var ctrlngValsObj,ctrlngValObj CR.ControllingVals
	var ctrldValsObj, ctrldValObj CR.ControlledVals
	var maxId int
	updateCtrlData = "Successfully created data"
	locColl.Find(bson.M{"zipcode":updateData.Loc.Zipcode, "country":updateData.Loc.Country}).Select(bson.M{"zipcode":1,"locid":1}).One(&locObj)

	if(len(locObj.Zipcode)==0){
		locColl.Find(bson.M{"locid":bson.M{"$ne":""}}).Select(
			bson.M{"locid":1}).Sort("-locid").Limit(1).One(&locObj)
		updateData.Loc.LocId = IDGen(locObj.LocId)

		dtimeColl.Find(bson.M{"conddateid":bson.M{"$ne":""}}).Select(
			bson.M{"conddateid":1}).Sort("-conddateid").Limit(1).One(&dtimeObj)
		updateData.Conddate.CondDateId  = IDGen(dtimeObj.CondDateId)
		updateData.Conddate.LocId = updateData.Loc.LocId

		ctrlgValsColl.Find(bson.M{"ctrlingvals.ctrlsvalsid":bson.M{"$ne":""}}).Select(
			bson.M{"ctrlingvals.ctrlvalsid":1}).Sort("-ctrlingvals.ctrlvalsid").Limit(1).One(&ctrlngValsObj)
		updateData.CtrlgVals.CtrlingVals[0].CtrlValsId = IDGen(ctrlngValsObj.CtrlingVals[0].CtrlValsId)
		updateData.CtrlgVals.CondDateId = updateData.Conddate.CondDateId

		ctrldValsColl.Find(bson.M{"ctrledvals.ctrldvalid":bson.M{"$ne":""}}).Select(
			bson.M{"ctrledvals.ctrldvalid":1}).Sort("-ctrledvals.ctrldvalid").Limit(1).One(&ctrldValsObj)
		updateData.CtrldVals.CtrledVals[0].CtrldValId = IDGen(ctrldValsObj.CtrledVals[0].CtrldValId)
		updateData.CtrldVals.CtrlValsId = updateData.CtrlgVals.CtrlingVals[0].CtrlValsId
		err = locColl.Insert(updateData.Loc)
		if err!= nil{
			return "failed creating data"
		}
		err = dtimeColl.Insert(updateData.Conddate)
		if err!= nil{
			return "failed creating data"
		}
		err = ctrlgValsColl.Insert(updateData.CtrlgVals)
		if err!= nil{
			return "failed creating data"
		}
		err = ctrldValsColl.Insert(updateData.CtrldVals)
		if err!= nil{
			return "failed creating data"
		}
	}else{
		updateData.Loc.LocId = locObj.LocId
		dtimeColl.Find(bson.M{"dtime":updateData.Conddate.Dtime, "working":updateData.Conddate.Working, "locid":locObj.LocId}).Select(
			bson.M{"dtime":1, "conddateid":1}).One(&dtimeObj )
		if(len(dtimeObj.Dtime)==0){
			dtimeColl.Find(bson.M{"conddateid":bson.M{"$ne":""}}).Select(
				bson.M{"conddateid":1}).Sort("-conddateid").Limit(1).One(&dtimeObj)
			updateData.Conddate.CondDateId  = IDGen(dtimeObj.CondDateId)
			updateData.Conddate.LocId = updateData.Loc.LocId

			ctrlgValsColl.Find(bson.M{"ctrlingvals.ctrlsvalsid":bson.M{"$ne":""}}).Select(
				bson.M{"ctrlingvals.ctrlvalsid":1}).Sort("-ctrlingvals.ctrlvalsid").Limit(1).One(&ctrlngValsObj)
			updateData.CtrlgVals.CtrlingVals[0].CtrlValsId = IDGen(ctrlngValsObj.CtrlingVals[0].CtrlValsId)
			updateData.CtrlgVals.CondDateId = updateData.Conddate.CondDateId

			ctrldValsColl.Find(bson.M{"ctrledvals.ctrldvalid":bson.M{"$ne":""}}).Select(
				bson.M{"ctrledvals.ctrldvalid":1}).Sort("-ctrledvals.ctrldvalid").Limit(1).One(&ctrldValsObj)
			updateData.CtrldVals.CtrledVals[0].CtrldValId = IDGen(ctrldValsObj.CtrledVals[0].CtrldValId)
			updateData.CtrldVals.CtrlValsId = updateData.CtrlgVals.CtrlingVals[0].CtrlValsId
			err = dtimeColl.Insert(updateData.Conddate)
			if err!= nil{
				return "failed creating data"
			}
			err = ctrlgValsColl.Insert(updateData.CtrlgVals)
			if err!= nil{
				return "failed creating data"
			}
			err = ctrldValsColl.Insert(updateData.CtrldVals)
			if err!= nil{
				return "failed creating data"
			}
		}else{
			updateData.Conddate.CondDateId = dtimeObj.CondDateId
			ctrlgValsColl.Find(bson.M{"conddateid":dtimeObj.CondDateId,
						  "ctrlingvals.condout":updateData.CtrlgVals.CtrlingVals[0].CondOut,
					          "ctrlingvals.tempout":updateData.CtrlgVals.CtrlingVals[0].TempOut,
				                  "ctrlingvals.lightout":updateData.CtrlgVals.CtrlingVals[0].LightOut,
						  "ctrlingvals.pplin":updateData.CtrlgVals.CtrlingVals[0].PplIn}).Select(
				bson.M{"ctrlingvals":1}).One(&ctrlngValsObj)
			if(ctrlngValsObj.CtrlingVals == nil){

				//check from here
				ctrlgValsColl.Find(bson.M{}).Select(
					bson.M{"ctrlingvals.ctrlvalsid":1}).Sort("-ctrlingvals.ctrlvalsid").Limit(1).One(&ctrlngValObj)
				maxId,_ = strconv.Atoi(ctrlngValObj.CtrlingVals[0].CtrlValsId)
				for _,j := range ctrlngValObj.CtrlingVals{
					thisId,_:= strconv.Atoi(j.CtrlValsId[3:])
					if(maxId < thisId && j.CtrlValsId != ""){
						maxId = thisId
					}
				}
				maxId += 1
				nextId := ctrlngValObj.CtrlingVals[0].CtrlValsId[:3]+strconv.Itoa(maxId)
				updateData.CtrlgVals.CtrlingVals[0].CtrlValsId = nextId
				updateData.CtrlgVals.CondDateId = updateData.Conddate.CondDateId
				ctrldValsColl.Find(bson.M{}).Select(
					bson.M{"ctrledvals.ctrldvalid":1}).Sort("-ctrledvals.ctrldvalid").Limit(1).One(&ctrldValObj)
				maxId = 0
				for _,j := range ctrldValObj.CtrledVals {
					thisId,_:= strconv.Atoi(j.CtrldValId[3:])
					if (maxId < thisId && j.CtrldValId != "") {
						maxId = thisId
					}
				}
				maxId += 1
				nextId = ctrldValObj.CtrledVals[0].CtrldValId[:3]+strconv.Itoa(maxId)
				updateData.CtrldVals.CtrledVals[0].CtrldValId = nextId
				updateData.CtrldVals.CtrlValsId = updateData.CtrlgVals.CtrlingVals[0].CtrlValsId

				err = ctrlgValsColl.Update(bson.M{"conddateid":updateData.Conddate.CondDateId},
					bson.M{"$push":bson.M{"ctrlingvals":updateData.CtrlgVals.CtrlingVals[0]}})
				if err!= nil{
					return "failed creating data"
				}
				err = ctrldValsColl.Insert(updateData.CtrldVals)
				if err!= nil{
					return "failed creating data"
				}
			}else{
				var controllingId string
				for _,j:= range ctrlngValsObj.CtrlingVals{
					if(j.CondOut == updateData.CtrlgVals.CtrlingVals[0].CondOut &&
						j.TempOut == updateData.CtrlgVals.CtrlingVals[0].TempOut &&
						j.LightOut == updateData.CtrlgVals.CtrlingVals[0].LightOut &&
						j.PplIn == updateData.CtrlgVals.CtrlingVals[0].PplIn){
						controllingId = j.CtrlValsId
						break
					}
				}

				updateData.CtrlgVals.CtrlingVals[0].CtrlValsId = controllingId
				ctrldValsColl.Find(bson.M{"ctrlvalsid":updateData.CtrlgVals.CtrlingVals[0].CtrlValsId,
							   "ctrledvals.tempin":updateData.CtrldVals.CtrledVals[0].TempIn,
							   "ctrledvals.lightin":updateData.CtrldVals.CtrledVals[0].LightIn,
							   "ctrledvals.musicin":updateData.CtrldVals.CtrledVals[0].MusicIn}).Select(
					bson.M{"ctrledvals":1}).One(&ctrldValsObj)
				if (ctrldValsObj.CtrledVals == nil) {
					ctrldValsColl.Find(bson.M{}).Select(
						bson.M{"ctrledvals.ctrldvalid":1}).Sort("-ctrledvals.ctrldvalid").Limit(1).One(&ctrldValObj)

					maxId = 0
					for _,j := range ctrldValObj.CtrledVals {
						thisId,_:= strconv.Atoi(j.CtrldValId[3:])
						if (maxId < thisId && j.CtrldValId != "") {
							maxId = thisId
						}
					}
					maxId += 1
					nextId := ctrldValObj.CtrledVals[0].CtrldValId[:3]+strconv.Itoa(maxId)
					updateData.CtrldVals.CtrledVals[0].CtrldValId = nextId
					err = ctrldValsColl.Update(bson.M{"ctrlvalsid":updateData.CtrlgVals.CtrlingVals[0].CtrlValsId},
						bson.M{"$push":bson.M{"ctrledvals":updateData.CtrldVals.CtrledVals[0]}})
					if err!= nil {
						return "failed creating data"
					}
				}else{
					for _,j:= range ctrldValsObj.CtrledVals{
						if(j.LightIn == updateData.CtrldVals.CtrledVals[0].LightIn &&
							j.TempIn == updateData.CtrldVals.CtrledVals[0].TempIn &&
							j.MusicIn == updateData.CtrldVals.CtrledVals[0].MusicIn ){

							err = ctrldValsColl.Update(bson.M{"ctrlvalsid":updateData.CtrlgVals.CtrlingVals[0].CtrlValsId,
								"ctrledvals.ctrldvalid":j.CtrldValId}, bson.M{"$set":bson.M{
								"ctrledvals.$.tempin":j.TempIn,"ctrledvals.$.lightin": j.LightIn,  "ctrledvals.$.musicin":j.MusicIn,
								"ctrledvals.$.homescount":j.HomesCount +1}})
							if err!= nil {
								panic(err)
								return "failed creating data"
							}
						}
					}
				}
			}

		}
	}

	return updateCtrlData
}
