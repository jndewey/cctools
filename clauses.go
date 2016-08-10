package cctools

//Creation of Clause struct and related methods; with JSON keys incorporated

type Clause struct {
	Name		string		`json:"name"`
	Text		string		`json:"text"`
	Description	string 		`json:"description"`
	Type 		string		`json:"type"`
	Benefit		bool		`json:"benefit"` //This mainly applies to conditions; and if true, means one party is benefitted by condition
	BftPty		Party		`json:"bftpty"` //If one or more parties benefit from a condition (to the exclusion of others), then BftPty identifies those parties
}

/* Getter for Name attribute

func (c Clause) Name() string {
	var output = c.name
	return output
}
*/

//Setter for Name attribute

func (c *Clause) SetName(n string) {
	c.Name = n
}

/* Getter for Text attribute

func (c Clause) Text() string {
	var output = c.Text
	return output
}
*/

//Setter for Text attribute

func (c *Clause) SetText(t string) {
	c.Text = t
}

/* Getter for Type attribute
//Initially supported types will be Definition, Covenant, Condition, 
//Representation (includes warranties) and Information (e.g., notice)

func (c Clause) Type() string {
	var output = c.Type
	return output
}
*/

//Setter for Type attribute

func (c *Clause) SetType(t string) {
	c.Type = t
}

/* Getter for Benefit attribute

func (c Clause) Benefit() bool {
	var output = c.Benefit
	return output
}
*/

//Setter for Benefit attribute

func (c *Clause) SetBenefit(b bool) {
	c.Benefit = b
}

/* Getter for BftPty attribute

func (c Clause) BftPty() Party {
	var output = c.BftPty
	return output
}
*/

//Setter for Benefit attribute
//Generally, the benefitted party has the right to waive a condition and proceed with the transaction

func (c *Clause) SetBftPty(b Party) {
	c.BftPty = b
}

//NOTE: This needs to be worked out because a changed clause will not be reflected in an existing documents because Clauses are a 
// type value rather than pointers to Clauses.  So the edited clause is not sharing data with the Clause that is an attribute
// in the Document object. 

func (c *Clause) EditText(t string) {
	c.Text = t
}


