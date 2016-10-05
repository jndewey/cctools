package cctools

import "gopkg.in/mgo.v2/bson"

//Creation of User struct and related methods

type User struct {
	Id           bson.ObjectId `bson:"_id,omitempty"` //By default, bson is used because our implementations use MondoDB.
	FirstName    string        `json:"firstName"`
	LastName     string        `json:"lastName"`
	DigitalID    string        `json:"digitalID"`
	Email        string        `json:"email"`
	PhoneNumber  string        `json:"phoneNumber"`
	AddressOne   string        `json:"addressOne"`
	AddressTwo   string        `json:"addressTwo"`
	AddressCity  string        `json:"adressCity"`
	AddressSt    string        `json:"addressSt"`
	AddressZip   string        `json:"addressZip"`
	Roles        []string      `json:"roles"`
	AccessRights []string      `json:"accessRights"`
}

func (u *User) SetId(n string) {
	u.Id = n
}

func (u *User) SetFirstName(n string) {
	u.FirstName = n
}

func (u *User) SetLastName(n string) {
	u.LastName = n
}

func (u *User) SetDigitalId(n string) {
	u.DigitalID = n
}

func (u *User) SetEmail(n string) {
	u.Email = n
}

func (u *User) SetPhoneNumber(n string) {
	u.PhoneNumber = n
}

func (u *User) SetAddressOne(n string) {
	u.AddressOne = n
}

func (u *User) SetAddressTwo(n string) {
	u.AddressTwo = n
}

func (u *User) SetAddressCity(n string) {
	u.AddressCity = n
}

func (u *User) SetAddressSt(n string) {
	u.AddressSt = n
}

func (u *User) SetAddressZip(n string) {
	u.AddressZip = n
}

func (u *User) AddRole(n string) {
	newRole := n
	oldRoles := u.Roles
	newRoles := append(oldRoles, newRole)
	u.Roles = newRoles
}

func (u *User) AddAccessRight(n string) {
	newAccessRight := n
	oldAccessRights := u.AccessRights
	newAccessRights := append(oldAccessRights, newAccessRight)
	u.AccessRights = newAccessRights
}
