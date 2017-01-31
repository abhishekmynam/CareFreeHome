package ConfigurationRepository

import "time"

type User struct {
	User_name string
	Pwd string
	UsrFName string
	UsrLName string
	Email string
	PwdCreate string
	PwdConfirm string
	Ph_no int64
	Address string
	City string
	Zipcode int64
	State string
	Country string
	UserStatus int
	CreatedDate time.Time
	LastUpdated time.Time
}

type Location struct{
	LocId string
	Zipcode string
	Country string
}

type CondDate struct{
	CondDateId string
	LocId string
	Dtime string
	Working bool
}

type ControllingVals struct{
	CondDateId string
	CtrlingVals []OutsideVals
}

type OutsideVals struct{
	CtrlValsId string
	CondOut string
	TempOut float64
	LightOut float64
	PplIn int64
}

type ControlledVals struct{
	CtrlValsId string
	CtrledVals []InsideVals
}

type InsideVals struct{
	CtrldValId string
	TempIn float64
	LightIn float64
	MusicIn float64
	HomesCount int64
}

type GlobalCtrlData struct{
	Loc Location
	Conddate CondDate
	CtrlgVals ControllingVals
	CtrldVals ControlledVals
}
