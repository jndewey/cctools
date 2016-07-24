package cctools

//Creation of Transaction struct including corresponding functions


type User struct {
	Id 			string
	FirstName 	string
	LastName	string
	DigitalID	string
	Email		string
	PhoneNumber	string
	AddressOne	string
	AddressTwo	string
	AddressCity	string
	AddressSte	string
	AddressZip	string
	Role 		[]string
	AccessRight []string
}

type Transaction struct {
	Id 			string
	Name 		string
	Client		string
	ClientNo	uint32
	Matter		string
	MatterNo	uint32 
	Documents 	[]Document
	Owner		User
	Users		[]User
	Checklist 	Checklist
}

type Checklist struct {
	Name 		string
	CheckItems	[]Item
	Documents 	[]Document
	Owner		string
	Users		[]User
}

type Item struct {
	Id 			string
	Name 		string
	Owner		string
	Users		[]User
	Done		bool
}

//Creation of Document struct, including corresponding functions

type Document struct {
	Id 			string
	FormNumber	int
	FornName	string
	Clauses		[]Clause
}

//Creation of Clause struct, including corresponding functions

type Clause struct {
	Name	string		`json:"name"`
	Text	string		`json:"text"`
}