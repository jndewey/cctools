package cctools

import "gopkg.in/mgo.v2/bson"

// import "time" UNCOMMENT WHEN DEADLINE COMPLETED

//Creation of Item struct and related methods

type Item struct {
	Id          bson.ObjectId `bson:"_id,omitempty"` //By default, bson is used because our implementations use MondoDB.
	Name        string        `json:"name"`
	Owner       string        `json:"owner"`
	Complete    bool          `json:"complete"`
	Category    bool          `json:"category"`
	Status      string        `json:"status"`
	PostClosing bool          `json:"postClosing"`
	//Add Deadline
}

/*
func NewItem(n string) Item {
	return Item{Id: "1", Name: n, Owner: "", Complete: false, Status: "Open", PostClosing: false}
}
*/
/*
func (i *Item) SetId(n string) {
	i.Id = n
}
*/
func (i *Item) SetName(n string) {
	i.Name = n
}

func (i *Item) SetOwner(s string) {
	i.Owner = s
}

/*
func (i *Item) AddUser(u User) {
	newUser := u
	oldUsers := i.Users
	newUsers := append(oldUsers, newUser)
	i.Users = newUsers
}
*/
func (i *Item) SetComplete(b bool) {
	i.Complete = b
}

func (i *Item) SetStatus(n string) {
	i.Status = n
}

func (i *Item) SetPostClosing(b bool) {
	i.PostClosing = b
}
