package cctools

//Creation of DyText struct and related methods; with JSON keys incorporated

type DyText struct {
	Name	string		`json:"name"`
	Params	[]string	`json:"params"`
	Type 	string		`json:"type"`
}

