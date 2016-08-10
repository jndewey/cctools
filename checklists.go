package cctools

//Creation of Checklist structure and related functions

type Checklist struct {
	Id	 		string		`json:"id"`
	Name 		string		`json:"name"`
	CheckItems	[]Item 		`json:"item"`
	Documents 	[]Document 	`json:"document"`
	Owner		string 		`json:"owner"`
	Users		[]User 		`json:"user"`
}

func (c *Checklist) SetId(n string) {
	c.Id = n
}

func (c *Checklist) SetName(n string) {
	c.Name = n
}

func (c *Checklist) AddItem(i Item) {
	newItem := i
	oldItems := c.CheckItems
	newItems := append(oldItems, newItem) 
	c.CheckItems = newItems
}

func (c *Checklist) AddDocument(d Document) {
	newDocument := d
	oldDocuments := c.Documents
	newDocuments := append(oldDocuments, newDocument) 
	c.Documents = newDocuments
}

func (c *Checklist) SetOwner(n string) {
	c.Owner = n
}

func (c *Checklist) AddUser(u User) {
	newUser := u
	oldUsers := c.Users
	newUsers := append(oldUsers, newUser) 
	c.Users = newUsers
}