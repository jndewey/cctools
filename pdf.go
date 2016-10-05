package cctools

import "github.com/jung-kurt/gofpdf"

// this will include method to convert Documents, Clauses, Checklists etc. to PDF format

func MakePDFClause(c Clause) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	text := c.Text
	name := c.Name
	pdf.Cell(40, 10, text)
	pdf.OutputFileAndClose(name + ".pdf")
}

func MakePDFDocument(d Document) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	clauses := d.Clauses
	var compiledDocument string
	compiledDocument = ""
	for _, clause := range clauses {
		compiledDocument = compiledDocument + clause.Text // FIX: Page content is built up using buffers (of type bytes.Buffer) rather than repeated string concatenation
	}
	text := compiledDocument
	name := d.Name
	pdf.Cell(40, 10, text)
	pdf.OutputFileAndClose(name + ".pdf")
}
