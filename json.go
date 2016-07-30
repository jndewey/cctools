package cctools

import "log"
import "encoding/json"
import "os"
import "fmt"
import "io/ioutil"

//JSON Conversion Functions for Individual Clauses

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

func JSONSave(data []byte) {
	file := data
	jsonFile, err := os.Create("./data.json") // replace with dynamic value

	if err !=nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(file)
	jsonFile.Close()
}

func LoadJSON(filename string) []byte {
	dat, err := ioutil.ReadFile(filename)
	if err !=nil {
		fmt.Println(err)
	}
	return dat
}

//JSON Conversion Functions for entire Documents

func DocumentToJSON(d Document) []byte {
	data, err := json.MarshalIndent(d, "", "    ")
	if err !=nil {
		log.Fatalf("JSON conversion failed: %s", err)
	}
	return data
}

func JSONtoDocument(data []byte) Document {
	var d Document
	json.Unmarshal(data, &d)
	return d
}

