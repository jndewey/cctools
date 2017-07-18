# CCTools
## contractCode Tools

ContractCode Tools is a Go library, which includes all of the necessary data structures and functions to create, edit and otherwise work with each of the following contractCode structures: Transactions, Documents and Clauses; all using the Go programming language.  In addition to these structures, cctools also exposes more granular building blocks like Users, Checklists, Items (as in checklist items) and Parties, among others.  This library is under active development, and changes almost daily, but the approach should become clear early in the development process.  While we had originally intended to build the entire contractCode platform in Javascript, it has become clear to us that Go is a great fit for the long term goals of the project, especially with respect to integration with blockchain protocols.  That said, an important principle guiding the development of cctools is modularity, both in terms of the structure of the library itself, but also in a broader web stack context.  When completed, we envision that cctools could plug into a Node.js, Ruby on Rails or just about any other stack as the model logic behind the contracts.  Our current plans include providing a basic front end UI for lawyers, but our primary focus will be on making cctools the perfect web service that feeds functionality (through standard HTTP APIs and/or SDKs) to applications with more robust front ends and database structures.

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

- Independent setter methods for all values. 

- Conversion Tools for JSON, YAML and XML, so that data structures, like a Document, can be converted to and from JSON, YAML and XML.  This is especially useful for storage of items in a database or other file systems.  For example, a record in a database that represents a Document can be retrieved and converted into a Document struct in Go.  In this form, a compliance module in the form of a Document method can be run against the Document (which is now in the form of a data structure capable of being passed to a method as a parameter).  In addition, data attributes of the Document (or any other data structure) can be modified before being converted back into a format suitable for storage in a database or other form of persistent storage.  

- The ability to easily convert from contractCode to structured data in the form of annotated text.  Properly annotated text is incredibly valuable for natural language processing and other machine learning tasks, but can be time consuming to produce.  ContractCode allows users to convert CC objects into XML formats with key features annotated based on parameter inputs. 

- The ability to convert text or html files into formatted DOCX or PDF files.  When it is necessary to express Documents, Checklists or other data structures in a traditional format, functions are provided that convert data structures into Word or PDF format.  As proof of the benefits of modularity, the tools for working with Word docx files are being programmed in Python in the form of Lambdas on Amazon Web Services.  This allows us to use the best language for this particular task, which happens to be Python, and expose the resulting functionality through simple HTTP calls to Lambda functions rather than setting up an entirely seperate web server stack.

## API and Documentation

APIs and documentation for all of the data structures and functionality are in process and will be published as and when completed.

## Installation and Code Example

In the interim, you can download this package into your GOPATH and import it into your Go programs.  The following code is an example of how to build clauses and combine them to form a document, and then how to convert back and forth between JSON and a native Go struct.
````go
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
````
## Future Development

We welcome any suggestions, proposed code or any other contributions that others want to make to the project.  We want cctools to develop into a robust tool that serves as the backbone to many other systems that allow a broad range of lawyers (from coding ninjas to total newbies) to leverage a more engineered approach to contract drafting (or as we like to call it, legal engineering).

## License

MIT License.

