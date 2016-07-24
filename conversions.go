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

func ReturnName(c Clause) string {
	var output = c.Name
	return output
}

func ReturnText(c Clause) string {
	var output = c.Text
	return output
}

func NumberOfClauses(d Document) int {
	return len(d.Clauses)
}

