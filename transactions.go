package cctools

//Creation of Transaction struct including corresponding functions

type Transaction struct {
	Id 			string
	Name 		string
	Documents 	[]Document
	Owner		string
	Users		[]string
	Private		bool

}