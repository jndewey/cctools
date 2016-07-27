package cctools

//Creation of User struct and related methods

type Party struct {
	Id 			string 		`json:"id"`
	Name 		string 		`json:"Name"`
	DefName		string		`json:"DefName"`
	Entity		bool		`json:"Entity"`
	Receiver	bool		`json:"Receiver"`
	StateOrg	string		`json:"StateOrg"`
	EntityType	string		`json:"EntityType"`
	DigitalID	string		`json:"digitalID"`
	Email		string 		`json:"email"`
	AddressOne	string 		`json:"addressOne"`
	AddressTwo	string 		`json:"addressTwo"`
	AddressCity	string 		`json:"adressCity"`
	AddressSte	string 		`json:"addressSte"`
	AddressZip	string 		`json:"addressZip"`
}