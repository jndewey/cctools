package cctools

//Creation of Transaction struct and related methods

type Transaction struct {
	Id 			string 		`json:"id"`
	Name 		string 		`json:"name"`
	Client		string 		`json:"client"`
	ClientNo	uint32 		`json:"clientNo"`
	Matter		string 		`json:"matter"`
	MatterNo	uint32 		`json:"matterNo"`
	Documents 	[]Document 	`json:"document"`
	Owner		User 		`json:"owner"`
	Users		[]User 		`json:"user"`
	Checklist 	Checklist 	`json:"checklist"`
	Marketing	string		`json:"marketing"`
}

