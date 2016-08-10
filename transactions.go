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
/*
func (t Transaction) Id() string {
	var output = t.Id
	return output
}
*/
func (t *Transaction) SetId(n string) {
	t.Id = n
}
/*
func (t Transaction) Name() string {
	var output = t.Name
	return output
}
*/
func (t *Transaction) SetName(n string) {
	t.Name = n
}
/*
func (t Transaction) Client() string {
	var output = t.Client
	return output
}
*/
func (t *Transaction) SetClient(n string) {
	t.Client = n
}
/*
func (t Transaction) ClientNo() uint32 {
	var output = t.ClientNo
	return output
}
*/
func (t *Transaction) SetClientNo(n uint32) {
	t.ClientNo = n
}
/*
func (t Transaction) Matter() string {
	var output = t.Matter
	return output
}
*/
func (t *Transaction) SetMatter(n string) {
	t.Matter = n
}
/*
func (t Transaction) MatterNo() uint32 {
	var output = t.MatterNo
	return output
}
*/
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
/*
func (t Transaction) Owner() User {
	var output = t.Owner
	return output
}
*/
func (t *Transaction) AddUser(u User) {
	newUser := u
	oldUsers := t.Users
	newUsers := append(oldUsers, newUser) 
	t.Users = newUsers
}

func (t *Transaction) SetChecklist(c Checklist) {
	t.Checklist = c
}
/*
func (t Transaction) Checklist() Checklist {
	var output = t.Checklist
	return output
}
*/
/*
func (t Transaction) Marketing() string {
	var output = t.Marketing
	return output
}
*/
func (t *Transaction) SetMarketing(n string) {
	t.Marketing = n
}
