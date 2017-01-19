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
	err = userColl.Insert(&User{usrFName: newUser.usrFName, usrLName:newUser.usrLName,
		email:newUser.email, pwdConfirm:newUser.pwdConfirm, ph_no: newUser.ph_no, address:newUser.address,
		city:newUser.city, zipcode:newUser.zipcode, state:newUser.state, country:newUser.country})
	if err!=nil{
		log.Fatal(err)
	}
}

