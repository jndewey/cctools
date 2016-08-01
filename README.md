# CCTools
contractCode Tools

ContractCode Tools is a Go library, which includes all of the necessary data structures and functions to create, edit and otherwise work with each of the following contractCode structures: Transactions, Contracts and Clauses; all using the Go programming language.  In addition to these structures, cctools also exposes more granular building blocks like users, checklist items and contract parties, among others.  This library is under active development, and changes almost daily, but the approach should become clear early in the development process.  While we had originally intended to build the entire contractCode platform in Javascript, it has become clear to us that Go is a great fit for the long term goals of the project, especially with respect to integration with blockchain protocols.  That said, an important principle guiding the development of cctools is modularity, both in terms of the structure of the library itself, but also in a broader web stack context.  When completed, we envision that cctools could plug into a Node.js, Ruby on Rails or just about any other stack as the model logic behind the contracts.

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


