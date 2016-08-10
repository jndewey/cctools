package cctools

//Creation of Document struct and related methods

type Document struct {
	Id 			string		`json:"id"`
	Name 		string		`json:"name"`
	Description	string		`json:"description"`
	FormNumber	int			`json:"formNumber"`
	FormName	string		`json:"formName"`
	Clauses		[]Clause 	`json:"clause"`
	Parties		[]Party		`json:"party"`
}

/*
func (d Document) Id() string {
	var output = d.Id
	return output
}
*/

func (d *Document) SetId(n string) {
	d.Id = n
}
/*
func (d Document) Name() string {
	var output = d.Name
	return output
}
*/
func (d *Document) SetName(n string) {
	d.Name = n
}
/*
func (d Document) Description() string {
	var output = d.Description
	return output
}
*/
func (d *Document) SetDescription(n string) {
	d.Description = n
}
/*
func (d Document) FormName() string {
	var output = d.FormName
	return output
}
*/
func (d *Document) SetFormName(n string) {
	d.FormName = n
}
/*
func (d Document) FormNumber() int {
	var output = d.FormNumber
	return output
}
*/
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

/* func (d *Document) DeleteClause(c Clause) {
	newClause := c
	oldClauses := d.Clauses
	newClauses := append(oldClauses, newClause) 
	d.Clauses = newClauses
} */
