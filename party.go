package cctools

//Creation of User struct and related methods

type Party struct {
	Id 			string 		`json:"id"`
	Name 		string 		`json:"name"`
	Definition	string		`json:"definition"`
	IsEntity	bool		`json:"isEntity"`
	IsReceiver	bool		`json:"isReceiver"`
	StateOrg	string		`json:"stateOrg"`
	EntityType	string		`json:"entityType"`
	DigitalID	string		`json:"digitalID"`
	Email		string 		`json:"email"`
	AddressOne	string 		`json:"addressOne"`
	AddressTwo	string 		`json:"addressTwo"`
	AddressCity	string 		`json:"adressCity"`
	AddressSt	string 		`json:"addressSt"`
	AddressZip	string 		`json:"addressZip"`
}

func (p *Party) SetId(n string) {
	p.Id = n
}

func (p *Party) SetName(n string) {
	p.Name = n
}

func (p *Party) SetDefinition(n string) {
	p.Definition = n
}

func (p *Party) SetIsEntity(b bool) {
	p.IsEntity = b
}

func (p *Party) SetIsReceiver(b bool) {
	p.IsReceiver = b
}

func (p *Party) SetStateOrg(n string) {
	p.StateOrg = n
}

func (p *Party) SetEntityType(n string) {
	p.EntityType = n
}

func (p *Party) SetDigitalId(n string) {
	p.DigitalID = n
}

func (p *Party) SetEmail(n string) {
	p.Email = n
}

func (p *Party) SetAddressOne(n string) {
	p.AddressOne = n
}

func (p *Party) SetAddressTwo(n string) {
	p.AddressTwo = n
}

func (p *Party) SetAddressCity(n string) {
	p.AddressCity = n
}

func (p *Party) SetAddressSt(n string) {
	p.AddressSt = n
}

func (p *Party) SetAddressZip(n string) {
	p.AddressZip = n
}
