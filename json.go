package cctools

import "log"
import "encoding/json"

//JSON Conversion Functions

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

