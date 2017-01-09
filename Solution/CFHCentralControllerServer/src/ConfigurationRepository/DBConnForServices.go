package ConfigurationRepository

import (
	"gopkg.in/mgo.v2"
	"log"
)

func SaveUser (newUser User ){
	session, err:= mgo.Dial("localhost:27017")
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	userColl := session.DB("CFHCentralDBTest").C("UserMaster")
	err = userColl.Insert(&User{newUser.usrFName, newUser.usrLName,
		newUser.email, newUser.pwdConfirm, newUser.ph_no, newUser.address,
		newUser.city, newUser.zipcode, newUser.state, newUser.country})
	if err!=nil{
		log.Fatal(err)
	}
}

