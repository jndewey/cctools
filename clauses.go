package cctools

//Creation of Clause struct and related methods; with JSON keys incorporated

type Clause struct {
	Name	string		`json:"name"`
	Text	string		`json:"text"`
	Type 	string		`json:"type"`
	Benefit	bool		`json:"benefit"` //This mainly applies to conditions; and if true, means one party is benefitted by condition
	BftPty	string		`json:"bftpty"` //If one party benefits from a condition, then BftPty identifies that party
}

//Getter for Name attribute

func Name(c Clause) string {
	var output = c.Name
	return output
}

//Setter for Name attribute

func (c *Clause) SetName(n string) {
	c.Name = n
}

//Getter for Text attribute

func Text(c Clause) string {
	var output = c.Text
	return output
}

//Setter for Text attribute

func (c *Clause) SetText(t string) {
	c.Text = t
}

//Getter for Type attribute
//Initially supported types will be definition, covenant, condition, 
//representation (includes warranties), information (e.g., notce)

func Type(c Clause) string {
	var output = c.Type
	return output
}

//Setter for Type attribute

func (c *Clause) SetType(t string) {
	c.Type = t
}

//Getter for Benefit attribute

func Benefit(c Clause) bool {
	var output = c.Benefit
	return output
}

//Setter for Benefit attribute

func (c *Clause) SetBenefit(b bool) {
	c.Benefit = b
}

//Getter for BftPty attribute

func BftPty(c Clause) string {
	var output = c.BftPty
	return output
}

//Setter for Benefit attribute

func (c *Clause) SetBftPty(b string) {
	c.BftPty = b
}

//NOTE: This needs to be worked out because a changed clause will not be reflected in an existing documents because Clauses are a 
// type value rather than pointers to Clauses.  So the edited clause is not sharing data with the Clause that is an attribute
// in the Document object. 

func (c *Clause) EditText(t string) {
	c.Text = t
}