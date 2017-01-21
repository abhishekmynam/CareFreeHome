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
