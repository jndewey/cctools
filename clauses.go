package cctools

//Creation of Clause struct and related methods

type Clause struct {
	Name	string		`json:"name"`
	Text	string		`json:"text"`
}

func Name(c Clause) string {
	var output = c.Name
	return output
}

func (c *Clause) SetName(n string) {
	c.Name = n
}

func Text(c Clause) string {
	var output = c.Text
	return output
}

func (c *Clause) SetText(t string) {
	c.Text = t
}


//NOTE: This needs to be worked out because a changed clause will not be reflected in an existing documents because Clauses are a 
// type value rather than pointers to Clauses.  So the edited clause is not sharing data with the Clause that is an attribute
// in the Document object. 

func (c *Clause) EditText(t string) {
	c.Text = t
}