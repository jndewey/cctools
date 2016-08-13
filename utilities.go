package cctools

import "html/template"
import "fmt"
import "bytes"


// Text Generators -- These tools convert Clauses and Documents into final text documents by combining parameter values with
// text templates for each clause or an entire Document.

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

// Compliance Utilities - Pass Documents and Clauses to method that takes a particular compliance module
// and applies the logic against the Document or Clause passed to the method

//Pass a Document to a compliance module to determine compliance

func (d *Document) CheckComplianceWith(m ComplianceModule) string {
	// this is just placeholder code + work in regex / parsing algorithms 
	output := "Now make it do something!" + " " + m.Name
	return output
}

//Pass a Clause to a compliance module to determine compliance

func (c *Clause) CheckComplianceWith(m ComplianceModule) string {
	// this is just placeholder code + work in regex / parsing algorithms 
	output := "Now make it do something!" + " " + c.Name
	return output
}

