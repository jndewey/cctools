package cctools

import "log"
import "encoding/json"
import "os"
import "fmt"
import "io/ioutil"
// import "github.com/ghodss/yaml" // REMOVE COMMENT WHEN FINISHED WITH YAML TOOLS


//Note that these YAML tools all use Go's standard library's JSON package for all functionality 
//except that JSON inputs and outputs are converted to and from YAML with ghodss/yaml package

//YAML Conversion Functions for Individual Clauses

func ClauseToYAML(c Clause) []byte {
	data, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		log.Fatalf("JSON conversion failed: %s", err)
	}
	return data
} 

func YAMLtoClause(data []byte) Clause {
	var c Clause
	json.Unmarshal(data, &c)
	return c
}

//Read and write individual YAML clause files

func SaveYAMLClause(data []byte, filename string) {
	file := data
	jsonFile, err := os.Create("./"+filename+".json")

	if err !=nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(file)
	jsonFile.Close()
}

func LoadYAMLClause(filename string) []byte {
	dat, err := ioutil.ReadFile(filename)
	if err !=nil {
		fmt.Println(err)
	}
	return dat
}

//YAML Conversion Functions for entire Documents

func DocumentToYAML(d Document) []byte {
	data, err := json.MarshalIndent(d, "", "    ")
	if err !=nil {
		log.Fatalf("JSON conversion failed: %s", err)
	}
	return data
}

func YAMLtoDocument(data []byte) Document {
	var d Document
	json.Unmarshal(data, &d)
	return d
}

//Read and write entire YAML Document files

func SaveYAMLDocument(data []byte, filename string) {
	file := data
	jsonFile, err := os.Create("./"+filename+".json")

	if err !=nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(file)
	jsonFile.Close()
}

func LoadYAMLDocument(filename string) []byte {
	dat, err := ioutil.ReadFile(filename)
	if err !=nil {
		fmt.Println(err)
	}
	return dat
}
