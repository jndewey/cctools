package cctools

//Creation of Transaction struct and related methods

type Transaction struct {
	Id 			string 		`json:"id"`
	Name 		string 		`json:"name"`
	Client		string 		`json:"client"`
	ClientNo	uint32 		`json:"clientNo"`
	Matter		string 		`json:"matter"`
	MatterNo	uint32 		`json:"matterNo"`
	Documents 	[]Document 	`json:"document"`
	Owner		User 		`json:"owner"`
	Users		[]User 		`json:"user"`
	Checklist 	Checklist 	`json:"checklist"`
	Marketing	string		`json:"marketing"`
}

func ImportTransaction(filename string) Transaction {
	data :=LoadFromJSON(filename)
	newTransaction :=JSONtoTransaction(data)
    return newTransaction
}

func (t *Transaction) SetId(n string) {
	t.Id = n
}

func (t *Transaction) SetName(n string) {
	t.Name = n
}

func (t *Transaction) SetClient(n string) {
	t.Client = n
}

func (t *Transaction) SetClientNo(n uint32) {
	t.ClientNo = n
}

func (t *Transaction) SetMatter(n string) {
	t.Matter = n
}

func (t *Transaction) SetMatterNo(n uint32) {
	t.MatterNo = n
}

func (t *Transaction) AddDocument(d Document) {
	newDocument := d
	oldDocuments := t.Documents
	newDocuments := append(oldDocuments, newDocument) 
	t.Documents = newDocuments
}

func (t *Transaction) SetOwner(u User) {
	t.Owner = u
}

func (t *Transaction) AddUser(u User) {
	newUser := u
	oldUsers := t.Users
	newUsers := append(oldUsers, newUser) 
	t.Users = newUsers
}

func (t *Transaction) SetChecklist(c Checklist) {
	t.Checklist = c
}

func (t *Transaction) SetMarketing(n string) {
	t.Marketing = n
}
