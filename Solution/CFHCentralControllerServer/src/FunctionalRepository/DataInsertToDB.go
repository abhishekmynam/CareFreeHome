package FunctionalRepository

import (
	"gopkg.in/mgo.v2"
	"log"
	CR "ConfigurationRepository"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func SaveUser (newUser CR.User ) string {
	var saveUserStat string
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	userColl := session.DB(CR.DBInstance).C(CR.UserMasterCollection)
	var existingUser CR.User
	err = userColl.Find(bson.M{"email":newUser.Email}).One(&existingUser)
	if (len(existingUser.Email)==0){
		err = userColl.Insert(&CR.User{UsrFName: newUser.UsrFName, UsrLName:newUser.UsrLName,
			Email:newUser.Email, PwdConfirm:newUser.PwdConfirm, Ph_no: newUser.Ph_no, Address:newUser.Address,
			City:newUser.City, Zipcode:newUser.Zipcode, State:newUser.State, Country:newUser.Country, UserStatus:newUser.UserStatus, CreatedDate:time.Now(),LastUpdated:time.Now()})
		saveUserStat ="New user created"
	}else if (len(existingUser.Email)==0){
		saveUserStat ="User already exists"
	}
	if err!=nil{
		log.Fatal(err)
		saveUserStat ="error occured check the logs"
	}
	return saveUserStat
}
/*
func SaveNewOutCondition (newUnProGenDataCol CR.UnProGenDataCol)string{
	var saveNewConditionOutStatus string
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	ctrlColl := session.DB(CR.DBInstance).C(CR.GlobalCtrlVals)
	var existingCondition CR.UnProGenDataCol
	err = ctrlColl.Find(bson.M{"condoutrecord":newUnProGenDataCol.CondOutRecord}).One(&existingCondition)
	if(len(existingCondition.CondOutRecord)==0){
		err = ctrlColl.Insert(newUnProGenDataCol)
	}
	return saveNewConditionOutStatus
}

*/