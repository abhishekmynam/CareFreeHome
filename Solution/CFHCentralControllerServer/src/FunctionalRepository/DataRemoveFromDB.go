package FunctionalRepository

import (
	"gopkg.in/mgo.v2"
	"log"
	CR "ConfigurationRepository"
	"gopkg.in/mgo.v2/bson"
)

func RemoveUser(email string) string{
	var removeUserStat string
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	userColl := session.DB(CR.DBInstance).C(CR.UserMasterCollection)
	var existingUser CR.User
	err = userColl.Find(bson.M{"email":email}).One(&existingUser)
	if (len(existingUser.Email)==0){
		removeUserStat = "User doesn't exist to remove"
	}else if(len(existingUser.Email)!=0){
		err = userColl.Remove(bson.M{"email":email})
		removeUserStat="User removed"
	}
	if err!=nil{
		log.Fatal(err)
		removeUserStat="error occured check the logs"
	}
	return removeUserStat
}
