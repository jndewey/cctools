package cctools

//Creation of Checklist structure and related functions

type Checklist struct {
	Name 		string		`json:"name"`
	CheckItems	[]Item 		`json:"item"`
	Documents 	[]Document 	`json:"document"`
	Owner		string 		`json:"owner"`
	Users		[]User 		`json:"user"`
}