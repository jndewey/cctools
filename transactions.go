package cctools

import "log"
import "encoding/json"

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
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
	id 			string
	formNumber	int
	fornName	string
	clauses		[]Clause
}

func NumberOfClauses(d Document) int {
	return len(d.clauses)
}

//Creation of Transaction struct including corresponding functions
type Transaction struct {
	id 			string
	name 		string
	documents 	[]Document
	owner		string
	users		[]string
	private		bool

}