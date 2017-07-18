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

//Read and write individual JSON clause files

func SaveJSONClause(data []byte, filename string) {
	file := data
	jsonFile, err := os.Create("./" + filename + ".json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(file)
	jsonFile.Close()
}

func LoadJSONClause(filename string) []byte {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return dat
}

//JSON Conversion Functions for entire Transactions

func TransactionToJSON(t Transaction) []byte {
	data, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		log.Fatalf("JSON conversion failed: %s", err)
	}
	return data
}

func JSONtoTransaction(data []byte) Transaction {
	var t Transaction
	json.Unmarshal(data, &t)
	return t
}

//JSON Conversion Functions for entire Documents

func DocumentToJSON(d Document) []byte {
	data, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		log.Fatalf("JSON conversion failed: %s", err)
	}
	return data
}

func JSONtoDocument(data []byte) Document {
	var d Document
	json.Unmarshal(data, &d)
	return d
}

//JSON Conversion Functions for entire Checklist

func ChecklistToJSON(c Checklist) []byte {
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		log.Fatalf("JSON conversion failed: %s", err)
	}
	return data
}

func JSONToChecklist(data []byte) Checklist {
	var c Checklist
	json.Unmarshal(data, &c)
	return c
}

//Read and write JSON files

func SaveToJSON(thing interface{}, filename string) {
	data, err := json.MarshalIndent(thing, "", "    ")
	if err != nil {
		log.Fatalf("JSON conversion failed: %s", err)
	}
	file := data
	jsonFile, err := os.Create("./" + filename + ".json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(file)
	jsonFile.Close()
}

func LoadFromJSON(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
