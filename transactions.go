package cctools

import "log"
import "encoding/json"


//Conversion Functions
//JSON
func ClauseToJSON(c Clause) []byte {
	data, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		log.Fatalf("JSON conversion failed: %s", err)
	}
	return data
} 

func JSONtoClause(data []byte) Clause {
	var c Clause
	json.Unmarshal(data, &c)
	return c
} 

//YAML
//TO BE ADDED

//Creation of Clause struct, including corresponding functions

type Clause struct {
	Name	string		`json:"name"`
	Text	string		`json:"text"`
}

func ReturnName(c Clause) string {
	var output = c.Name
	return output
}

func ReturnText(c Clause) string {
	var output = c.Text
	return output
}

//Creation of Document struct, including corresponding functions

type Document struct {
	Id 			string
	FormNumber	int
	FornName	string
	Clauses		[]Clause
}

func NumberOfClauses(d Document) int {
	return len(d.Clauses)
}

//Creation of Transaction struct including corresponding functions
type Transaction struct {
	Id 			string
	Name 		string
	Documents 	[]Document
	Owner		string
	Users		[]string
	Private		bool

}