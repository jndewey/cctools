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