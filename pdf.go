package cctools

import "github.com/jung-kurt/gofpdf"

// this will include method to convert Documents, Clauses, Checklists etc. to PDF format

func MakePDF(c Clause) {
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.AddPage()
pdf.SetFont("Arial", "B", 16)
text := c.Text
name := c.Name
pdf.Cell(40, 10, text)
pdf.OutputFileAndClose(name + ".pdf")

}