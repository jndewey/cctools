package cctools

import "gopkg.in/mgo.v2/bson"

//The Checklist struct serves as the backbone of the transaction management component of contractCode.
//In addition to a unique id, a Checklist has a name (e.g., ABC Bank Loan to ACME Corp.), a slice of Items,
//an Owner, who is generally the creator of the checklist and a slice of Users, who can each have roles
//in relation to the Checklist.
type Checklist struct {
	Id         bson.ObjectId `bson:"_id,omitempty"` //By default, bson is used because our implementations use MondoDB.
	Name       string        `json:"name"`
	CheckItems []Item        `json:"item"`
	Owner      User          `json:"owner"`
	Users      []User        `json:"user"`
}

func ImportChecklist(filename string) Checklist {
	data := LoadFromJSON(filename)
	newChecklist := JSONToChecklist(data)
	return newChecklist
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
