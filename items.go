package cctools

//Creation of Item struct and related methods

type Item struct {
	Id 			string		`json:"id"`
	Name 		string		`json:"name"`
	Owner		string		`json:"owner"`
	Users		[]User 		`json:"users"`	
	Done		bool 		`json:"done"`
}