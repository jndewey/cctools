package cctools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
)

//Creation of Clause struct and related methods; with JSON keys incorporated
type Clause struct {
	Name        string            `json:"name"`
	Heading     string            `json:"heading"`
	Text        string            `json:"text"`
	Description string            `json:"description"`
	Tags        []string          `json:"tags"`
	Type        string            `json:"type"`
	Params      map[string]string `json:"params"`
	MergedText  string            `json:"mergedtext"`
}

func ImportClause(filename string) Clause {
	data := LoadFromJSON(filename)
	newClause := JSONtoClause(data)
	return newClause
}

// Converts Params field of Clause struct into JSON object for delivery to client application
func (c *Clause) ParamstoJSON() []byte {
	paramsJSON, _ := json.Marshal(c.Params)
	return paramsJSON
}

//Setter for Name attribute
func (c *Clause) SetName(n string) {
	c.Name = n
}

//Setter for Heading attribute
func (c *Clause) SetHeading(h string) {
	c.Heading = h
}

//Setter for Text attribute
func (c *Clause) SetText(t string) {
	c.Text = t
}

//Setter for Description attribute
func (c *Clause) SetDescription(d string) {
	c.Description = d
}

func (c *Clause) AddTag(t string) {
	newTag := t
	oldTags := c.Tags
	newTags := append(oldTags, newTag)
	c.Tags = newTags
}

//Setter for Type attribute
func (c *Clause) SetType(t string) {
	c.Type = t
}

//This function accepts a pointer to a Clause and renders the template text by passing the Parameter values to the template text
//and afterwards saves the resulting merged text to the MergedText attribute of the Clause
func (c *Clause) ParseClause() {
	t, err := template.New(c.Name).Parse(c.Text)
	if err != nil {
		fmt.Println(err)
	}
	varmap := c.Params
	data := new(bytes.Buffer)
	t.ExecuteTemplate(data, "TextForTemplate", varmap)
	output := data.String()
	c.MergedText = output
}

//NOTE: This needs to be worked out because a changed clause will not be reflected in an existing documents because Clauses are a
// type value rather than pointers to Clauses.  So the edited clause is not sharing data with the Clause that is an attribute
// in the Document object.
func (c *Clause) EditText(t string) {
	c.Text = t
}
