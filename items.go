package cctools

// import "time" UNCOMMENT WHEN DEADLINE COMPLETED

//Creation of Item struct and related methods

type Item struct {
	Id 			string		`json:"id"`
	Name 		string		`json:"name"`
	Owner		User		`json:"owner"`
	Users		[]User 		`json:"users"`	
	Complete	bool 		`json:"complete"`
	Status		string		`json:"status"`
	PostClosing	bool		`json:"postClosing"`
	//Add Deadline
}

func (i *Item) SetId(n string) {
	i.Id = n
}

func (i *Item) SetName(n string) {
	i.Name = n
}

func (i *Item) SetOwner(u User) {
	i.Owner = u
}

func (i *Item) AddUser(u User) {
	newUser := u
	oldUsers := i.Users
	newUsers := append(oldUsers, newUser) 
	i.Users = newUsers
}

func (i *Item) SetComplete(b bool) {
	i.Complete = b
}

func (i *Item) SetStatus(n string) {
	i.Status = n
}

func (i *Item) SetPostClosing(b bool) {
	i.PostClosing = b
}