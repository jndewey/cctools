package cctools

// Compliance Utilities - Pass Documents and Clauses to method that takes a particular compliance module
// and applies the logic against the Document or Clause passed to the method

//Pass a Document to a compliance module to determine compliance

func (d *Document) CheckComplianceWith(m ComplianceModule) string {
	// this is just placeholder code + work in regex / parsing algorithms 
	output := "Now make it do something!" + " " + m.Name
	return output
}

//Pass a Clause to a compliance module to determine compliance

func (c *Clause) CheckComplianceWith(m ComplianceModule) string {
	// this is just placeholder code + work in regex / parsing algorithms 
	output := "Now make it do something!" + " " + c.Name
	return output
}



/* MakeTransaction function creates and initializes a new transaction based on a transaction template

func MakeTransaction(tmp TransactionTemp, name string) Transaction {
	newTransaction := tmp
	newTransaction.Name = name
	return newTransaction
}
*/