package cctools

//Creation of Checklist structure and related functions

type Checklist struct {
	Id	 		string		`json:"id"`
	Name 		string		`json:"name"`
	CheckItems	[]Item 		`json:"item"`
	Owner		User 		`json:"owner"`
	Users		[]User 		`json:"user"`
}

func ImportChecklist(filename string) Checklist {
	data :=LoadFromJSON(filename)
	newChecklist :=JSONtoChecklist(data)
    return  newChecklist
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

func (c *Checklist) SetOwner(u User) {
	c.Owner = u
}

func (c *Checklist) AddUser(u User) {
	newUser := u
	oldUsers := c.Users
	newUsers := append(oldUsers, newUser) 
	c.Users = newUsers
}