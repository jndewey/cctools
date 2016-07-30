# CCTools
contractCode Tools

ContractCode Tools is a Go library, which includes all of the necessary data structures and functions to create, edit and otherwise work with each of the following contractCode structures: Transactions, Contracts and Clauses; all using the Go programming language.  In addition to each of the foregoing structures, this library also exposes the building blocks for these items (e.g., users, checklist items and contract parties).  

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


