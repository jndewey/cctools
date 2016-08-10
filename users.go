package cctools

//Creation of User struct and related methods

type User struct {
	Id 			string 		`json:"id"`
	FirstName 	string 		`json:"firstName"`
	LastName	string 		`json:"lastName"`
	DigitalID	string		`json:"digitalID"`
	Email		string 		`json:"email"`
	PhoneNumber	string 		`json:"phoneNumber"`
	AddressOne	string 		`json:"addressOne"`
	AddressTwo	string 		`json:"addressTwo"`
	AddressCity	string 		`json:"adressCity"`
	AddressSte	string 		`json:"addressSte"`
	AddressZip	string 		`json:"addressZip"`
	Role 		[]string 	`json:"role"`
	AccessRight []string 	`json:"accessRight"`
}

//add methods