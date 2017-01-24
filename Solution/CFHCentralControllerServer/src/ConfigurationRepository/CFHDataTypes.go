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

type ControllingVals struct{
	TempInside float64
	LightsInside float64
	MusicInside float64
	CountOfhomes int64
}

type ControllingPpl struct {
	PplCount int64
	InsideControlVal ControllingVals
}

type ControllingMusic struct {
	MusicLevel float64
	InsideControlPpl ControllingPpl
}

type ControllingLights struct {
	OutsideLight float64
	InsideControlMusic ControllingMusic
}

type ControllingTemp struct{
	OutsideTemp float64
	OutsideControlLight ControllingLights
}

type AreaZip struct{
	ZipCode int64
	OutsideControlTemp ControllingTemp
}

type RecordTime struct {
	TimeRecord float64
	AreaZipToRecord AreaZip
}

type RecordDateOfMonth struct {
	DateRecord int64
	ControlTime RecordTime
}

type RecordMonth struct{
	MonthRecord string
	ControlDate RecordDateOfMonth
}

type UnProGenDataCol struct{
	CondOutRecord string
	ControlMonth RecordMonth
}

