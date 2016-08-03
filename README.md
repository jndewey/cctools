# CCTools
contractCode Tools

ContractCode Tools is a Go library, which includes all of the necessary data structures and functions to create, edit and otherwise work with each of the following contractCode structures: Transactions, Contracts and Clauses; all using the Go programming language.  In addition to these structures, cctools also exposes more granular building blocks like users, checklist items and contract parties, among others.  This library is under active development, and changes almost daily, but the approach should become clear early in the development process.  While we had originally intended to build the entire contractCode platform in Javascript, it has become clear to us that Go is a great fit for the long term goals of the project, especially with respect to integration with blockchain protocols.  That said, an important principle guiding the development of cctools is modularity, both in terms of the structure of the library itself, but also in a broader web stack context.  When completed, we envision that cctools could plug into a Node.js, Ruby on Rails or just about any other stack as the model logic behind the contracts.  Our current plans include providing a basic front end UI for lawyers, but our primary focus will be on making cctools the perfect web service that feeds functionality (through standard HTTP APIs) to applications with more robust front ends and database structures.

The contractCode White Paper can be found here: https://docs.google.com/document/d/1Cun8B6V_CbedxrhW26j0ZfAfcuVKtrVOdg9tY7XR8Lw/edit?usp=sharing

More specifically, the following structs, interfaces and other data structures are available:

Types:

User -- Someone with rights to access deal information.

Party -- A party to a contract.

Transaction -- A deal.

Checklist -- A checklist for a transaction.

Item -- A checklist item.

Document -- Most commonly used as a contract, but can be used more broadly.

Clause -- A clause that can be included within a contract (or Document).

In addition to the above types, each of the following functions and methods are available:

Setters and getters for all values. 

Conversion Tools for JSON, YAML and XML, so that data structures, like a Document, can be converted to and from JSON etc. This is especially useful for storage of items in a database or other file system.

Ability to convert text or html files into formatted DOCX or PDF files.  For traditional contracting, the data structures will need to be capable of conversion to Word or PDF.

Detailed Documentation is in process.

In the interim, you can download this package into your GOPATH and import it into your Go programs.  The following code is an example of how to build a short promissory note.

package main 

import "fmt"
import "github.com/jndewey/cctools" // replace this import path with the correct path to cctools on your system

func main() {
	//Create a couple of sample clauses
	var promisePay cctools.Clause
	promisePay.Name = "Promise to Pay"
	promisePay.Text = "Maker hereby promises to pay to the order of Payee the Principal Amount, together with Accrued Interest, in accordance with the provisions of this Promissory Note."
	var governingLaw cctools.Clause
	governingLaw.Name = "Governing Law"
	governingLaw.Text = "This Promissory Note shall be governed by the laws of the State of Florida."
	var venue cctools.Clause
	venue.Name = "Venue"
	venue.Text = "Venue will be in Miami-Dade County, Florida"
	//Now create a document and add your clauses
	var contractTest cctools.Document
	contractTest.Clauses = []cctools.Clause {promisePay, venue, governingLaw}
	//We can now convert that contract to JSON format if we want
	var converted = cctools.DocumentToJSON(contractTest)
	fmt.Printf("%s\n", converted)
	//We can also take JSON format and convert to a Clause struct in Go
	jdata := []byte(`{"name": "Venue", "text": "Venue will be in Palm Beach County"}`)
	var structOutput = cctools.JSONtoClause(jdata)
	fmt.Println(structOutput.Text)
	//We can also determine the number of clauses in a Document with the Length method
	num := cctools.Length(contractTest)
	fmt.Println(num)
}

