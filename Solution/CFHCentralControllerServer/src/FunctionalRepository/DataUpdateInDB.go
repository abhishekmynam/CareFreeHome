package FunctionalRepository

import (
	"gopkg.in/mgo.v2"
	"log"
	CR "ConfigurationRepository"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func UpdateUser (newUser CR.User) string {
	var DeactivateUserStatus string
	session, err := mgo.Dial(CR.DBServerTest)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	userColl := session.DB(CR.DBInstanceTest).C(CR.UserMasterCollection)
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
