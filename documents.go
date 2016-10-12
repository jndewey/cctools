package cctools

import "gopkg.in/mgo.v2/bson"

//Creation of Document struct and related methods

type Document struct {
	Id          bson.ObjectId `bson:"_id,omitempty"` //By default, bson is used because our implementations use MondoDB.
	Name        string        `json:"name"`
	Description string        `json:"description"`
	FormNumber  int           `json:"formNumber"`
	FormName    string        `json:"formName"`
	Clauses     []Clause      `json:"clause"`
	Parties     []Party       `json:"party"`
	Priority    int           `json:"priority"`
}

func ImportDocument(filename string) Document {
	data := LoadFromJSON(filename)
	newDocument := JSONtoDocument(data)
	return newDocument
}

/*
func (d *Document) SetId(n string) {
	d.Id = n
}
*/
func (d *Document) SetName(n string) {
	d.Name = n
}

func (d *Document) SetDescription(n string) {
	d.Description = n
}

func (d *Document) SetFormName(n string) {
	d.FormName = n
}

func (d *Document) SetFormNumber(n int) {
	d.FormNumber = n
}

func Length(d Document) int {
	return len(d.Clauses)
}

func (d *Document) AddClause(c Clause) {
	newClause := c
	oldClauses := d.Clauses
	newClauses := append(oldClauses, newClause)
	d.Clauses = newClauses
}
