package FunctionalRepository

import (
	"gopkg.in/mgo.v2"
	"log"
	CR "ConfigurationRepository"
	"gopkg.in/mgo.v2/bson"
	"time"
	//"fmt"
)

func UpdateUser1 (newUser CR.User) string {
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
/*
func UpdateGenControlData (updateData CR.UnProGenDataCol)string{
	var updateCtrlData string
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	ctrlColl := session.DB(CR.DBInstance).C(CR.GlobalCtrlVals)

	var existingCond CR.UnProGenDataCol
	ctrlColl.Find(bson.M{"condoutrecord":updateData.CondOutRecord}).Select(bson.M{"condoutrecord":1}).One(&existingCond)
	if(len(existingCond.CondOutRecord)==0){
		err = ctrlColl.Insert(updateData)
	}else {
		pipeline := []bson.M{
			{"$match":bson.M{"condoutrecord":updateData.CondOutRecord}},
			//{"$unwind":"$controlmonth"},
			{"$match":bson.M{"controlmonth.monthrecord":updateData.ControlMonth[0].MonthRecord}, },
			{"$project":bson.M{"condoutrecord":1}},
		}
		var existingMon CR.UnProGenDataCol
		ctrlColl.Pipe(pipeline).One(&existingMon)

		if (len(existingMon.CondOutRecord) == 0) {

			err = ctrlColl.Update(bson.M{"condoutrecord":updateData.CondOutRecord},
				bson.M{"$push":bson.M{"controlmonth": updateData.ControlMonth[0]}})//bson.M{"$monthrecord":updateData.ControlMonth[0].MonthRecord,"$controldate":updateData.ControlMonth[0].ControlDate}}})
		} else {
			pipeline := []bson.M{
				{"$match":bson.M{"condoutrecord":updateData.CondOutRecord}},
				{"$match":bson.M{"controlmonth.monthrecord":updateData.ControlMonth[0].MonthRecord}},
				{"$match":bson.M{"controlmonth.controldate.daterecord":updateData.ControlMonth[0].ControlDate[0].DateRecord}},
				{"$project":bson.M{"condoutrecord":1}},
			}
			var existingDate CR.UnProGenDataCol
			ctrlColl.Pipe(pipeline).One(&existingDate)

			if (len(existingDate.CondOutRecord) == 0) {
				err = ctrlColl.Update(bson.M{"condoutrecord":updateData.CondOutRecord,
							     "controlmonth.monthrecord":updateData.ControlMonth[0].MonthRecord},
						      bson.M{"$push":bson.M{"controlmonth.$.controldate": updateData.ControlMonth[0].ControlDate[0]}})
			} else {
				pipeline := []bson.M{
					{"$match":bson.M{"condoutrecord":updateData.CondOutRecord}},
					{"$match":bson.M{"controlmonth.monthrecord":updateData.ControlMonth[0].MonthRecord}},
					{"$match":bson.M{"controlmonth.controldate.daterecord":updateData.ControlMonth[0].ControlDate[0].DateRecord}},
					{"$match":bson.M{"controlmonth.controldate.controltime.timerecord":updateData.ControlMonth[0].ControlDate[0].ControlTime[0].TimeRecord}},
					//{"$project":bson.M{"condoutrecord":1}},
				}
				var existingTime CR.UnProGenDataCol
				ctrlColl.Pipe(pipeline).One(&existingTime)
				if(len(existingTime.CondOutRecord) == 0){
					err = ctrlColl.Update(bson.M{"condoutrecord":updateData.CondOutRecord,
						                     "controlmonth.monthrecord":updateData.ControlMonth[0].MonthRecord,
						                     "controlmonth.controldate.daterecord":updateData.ControlMonth[0].ControlDate[0].DateRecord},
							bson.M{"$push":bson.M{"controlmonth.controldate.0.controltime": updateData.ControlMonth[0].ControlDate[0].ControlTime[0]}})
				}
				fmt.Println("thsi")
			}
		}
	/*

				else{
					pipeline := []bson.M{
						{"$match":bson.M{"condoutrecord":updateData.CondOutRecord}},
						{"$unwind":"$controlmonth"},
						{"$match":bson.M{"controlmonth.monthrecord":updateData.ControlMonth.MonthRecord}},
						{"$unwind":"$controlmonth.controldate"},
						{"$match":bson.M{"controlmonth.controldate.daterecord":updateData.ControlMonth.ControlDate.DateRecord}},
						{"$unwind":"$controlmonth.controldate.controltime"},
						{"$match":bson.M{"controlmonth.controldate.controltime.timerecord":updateData.ControlMonth.ControlDate.ControlTime.TimeRecord}},
						{"$unwind":"$controlmonth.controldate.controltime"},
						{"$match":bson.M{"controlmonth.controldate.controltime.areaziptorecord.zipcode":updateData.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.ZipCode}},
						{"$project":bson.M{"controlmonth.controldate.controltime.areaziptorecord.zipcode":1}},
					}
					var existingZip CR.UnProGenDataCol
					existingZip.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.ZipCode = -999
					ctrlColl.Pipe(pipeline).One(&existingZip)

					if(existingZip.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.ZipCode == -999){

					}else{
						pipeline := []bson.M{
							{"$match":bson.M{"condoutrecord":updateData.CondOutRecord}},
							{"$unwind":"$controlmonth"},
							{"$match":bson.M{"controlmonth.monthrecord":updateData.ControlMonth.MonthRecord}},
							{"$unwind":"$controlmonth.controldate"},
							{"$match":bson.M{"controlmonth.controldate.daterecord":updateData.ControlMonth.ControlDate.DateRecord}},
							{"$unwind":"$controlmonth.controldate.controltime"},
							{"$match":bson.M{"controlmonth.controldate.controltime.timerecord":updateData.ControlMonth.ControlDate.ControlTime.TimeRecord}},
							{"$unwind":"$controlmonth.controldate.controltime.areaziptorecord"},
							{"$match":bson.M{"controlmonth.controldate.controltime.areaziptorecord.zipcode":updateData.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.ZipCode}},
							{"$unwind":"$controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp"},
							{"$match":bson.M{"controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidetemp":updateData.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.OutsideControlTemp.OutsideTemp}},
							{"$project":bson.M{"controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidetemp":1}},
						}
						var existingOutTemp CR.UnProGenDataCol
						existingOutTemp.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.OutsideControlTemp.OutsideTemp= -999
						ctrlColl.Pipe(pipeline).One(&existingOutTemp)

						if(existingOutTemp.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.OutsideControlTemp.OutsideTemp==-999){

						}else{
							pipeline := []bson.M{
								{"$match":bson.M{"condoutrecord":updateData.CondOutRecord}},
								{"$unwind":"$controlmonth"},
								{"$match":bson.M{"controlmonth.monthrecord":updateData.ControlMonth.MonthRecord}},
								{"$unwind":"$controlmonth.controldate"},
								{"$match":bson.M{"controlmonth.controldate.daterecord":updateData.ControlMonth.ControlDate.DateRecord}},
								{"$unwind":"$controlmonth.controldate.controltime"},
								{"$match":bson.M{"controlmonth.controldate.controltime.timerecord":updateData.ControlMonth.ControlDate.ControlTime.TimeRecord}},
								{"$unwind":"$controlmonth.controldate.controltime.areaziptorecord"},
								{"$match":bson.M{"controlmonth.controldate.controltime.areaziptorecord.zipcode":updateData.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.ZipCode}},
								{"$unwind":"$controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp"},
								{"$match":bson.M{"controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidetemp":updateData.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.OutsideControlTemp.OutsideTemp}},
								{"$unwind":"$controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidecontrollight"},
								{"$match":bson.M{"controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidecontrollight.outsidelight":updateData.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.OutsideControlTemp.OutsideControlLight.OutsideLight}},
								{"$project":bson.M{"controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidecontrollight.outsidelight":1}},
							}
							var existingOutLight CR.UnProGenDataCol
							existingOutLight.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.OutsideControlTemp.OutsideControlLight.OutsideLight= -999
							ctrlColl.Pipe(pipeline).One(&existingOutLight)

							if(existingOutLight.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.OutsideControlTemp.OutsideControlLight.OutsideLight== -999){

							}else{
								pipeline := []bson.M{
									{"$match":bson.M{"condoutrecord":updateData.CondOutRecord}},
									{"$unwind":"$controlmonth"},
									{"$match":bson.M{"controlmonth.monthrecord":updateData.ControlMonth.MonthRecord}},
									{"$unwind":"$controlmonth.controldate"},
									{"$match":bson.M{"controlmonth.controldate.daterecord":updateData.ControlMonth.ControlDate.DateRecord}},
									{"$unwind":"$controlmonth.controldate.controltime"},
									{"$match":bson.M{"controlmonth.controldate.controltime.timerecord":updateData.ControlMonth.ControlDate.ControlTime.TimeRecord}},
									{"$unwind":"$controlmonth.controldate.controltime.areaziptorecord"},
									{"$match":bson.M{"controlmonth.controldate.controltime.areaziptorecord.zipcode":updateData.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.ZipCode}},
									{"$unwind":"$controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp"},
									{"$match":bson.M{"controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidetemp":updateData.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.OutsideControlTemp.OutsideTemp}},
									{"$unwind":"$controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidecontrollight"},
									{"$match":bson.M{"controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidecontrollight.outsidelight":updateData.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.OutsideControlTemp.OutsideControlLight.OutsideLight}},
									{"$unwind":"$controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidecontrollight.insidecontrolmusic"},
									{"$match":bson.M{"controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidecontrollight.insidecontrolmusic.musiclevel":updateData.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.OutsideControlTemp.OutsideControlLight.InsideControlMusic.MusicLevel}},
									{"$project":bson.M{"controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidecontrollight.insidecontrolmusic.musiclevel":1}},
								}
								var existingInMusic CR.UnProGenDataCol
								existingInMusic.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.OutsideControlTemp.OutsideControlLight.InsideControlMusic.MusicLevel= -999
								ctrlColl.Pipe(pipeline).One(&existingInMusic)

								if (existingInMusic.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.OutsideControlTemp.OutsideControlLight.InsideControlMusic.MusicLevel== -999){

								}else{
									pipeline := []bson.M{
										{"$match":bson.M{"condoutrecord":updateData.CondOutRecord}},
										{"$unwind":"$controlmonth"},
										{"$match":bson.M{"controlmonth.monthrecord":updateData.ControlMonth.MonthRecord}},
										{"$unwind":"$controlmonth.controldate"},
										{"$match":bson.M{"controlmonth.controldate.daterecord":updateData.ControlMonth.ControlDate.DateRecord}},
										{"$unwind":"$controlmonth.controldate.controltime"},
										{"$match":bson.M{"controlmonth.controldate.controltime.timerecord":updateData.ControlMonth.ControlDate.ControlTime.TimeRecord}},
										{"$unwind":"$controlmonth.controldate.controltime.areaziptorecord"},
										{"$match":bson.M{"controlmonth.controldate.controltime.areaziptorecord.zipcode":updateData.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.ZipCode}},
										{"$unwind":"$controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp"},
										{"$match":bson.M{"controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidetemp":updateData.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.OutsideControlTemp.OutsideTemp}},
										{"$unwind":"$controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidecontrollight"},
										{"$match":bson.M{"controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidecontrollight.outsidelight":updateData.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.OutsideControlTemp.OutsideControlLight.OutsideLight}},
										{"$unwind":"$controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidecontrollight.insidecontrolmusic"},
										{"$match":bson.M{"controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidecontrollight.insidecontrolmusic.musiclevel":updateData.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.OutsideControlTemp.OutsideControlLight.InsideControlMusic.MusicLevel}},
										{"$unwind":"$controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidecontrollight.insidecontrolmusic.insidecontrolppl"},
										{"$match":bson.M{"controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidecontrollight.insidecontrolmusic.insidecontrolppl.pplcount":updateData.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.OutsideControlTemp.OutsideControlLight.InsideControlMusic.InsideControlPpl.PplCount}},
										{"$project":bson.M{"controlmonth.controldate.controltime.areaziptorecord.outsidecontroltemp.outsidecontrollight.insidecontrolmusic.insidecontrolppl.pplcount":1}},
									}

									var existingInPpl CR.UnProGenDataCol
									existingInPpl.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.OutsideControlTemp.OutsideControlLight.InsideControlMusic.InsideControlPpl.PplCount= -999
									ctrlColl.Pipe(pipeline).One(&existingInPpl)

									if(existingInPpl.ControlMonth.ControlDate.ControlTime.AreaZipToRecord.OutsideControlTemp.OutsideControlLight.InsideControlMusic.InsideControlPpl.PplCount== -999){

									}else{

									}
									fmt.Println(existingInMusic)
								}
							}
						}
					}
				}
			}

		}
	*/
	/*}
	if(err!=nil){
		panic(err)
	}


	return updateCtrlData
}
*/