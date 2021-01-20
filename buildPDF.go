package main

import "github.com/jung-kurt/gofpdf"

// PDFDoc struct data
type PDFDoc struct {
	document     *gofpdf.Fpdf
	encodingFunc func(string) string
}

var fontPtSize float64 = 12

// NewPdfDoc function displays results
func NewPdfDoc() *PDFDoc {
	pdf := gofpdf.New("P", "mm", "A4", "./assets/fonts")
	// html := pdf.HTMLBasicNew()
	encodingFunc := pdf.UnicodeTranslatorFromDescriptor("")

	pdf.AddFont("ProximaNova", "", "ProximaNova-Reg-webfont.json")
	pdf.AddFont("ProximaNova", "B", "ProximaNova-Bold-webfont.json")
	pdf.AddFont("ProximaNova-Light", "", "ProximaNova-Light-webfont.json")

	doc := &PDFDoc{
		// htmlContent:  html,
		encodingFunc: encodingFunc,
	}

	pdf.SetFont("ProximaNova", "", fontPtSize)
	pdf.SetTextColor(75, 75, 80)
	pdf.AliasNbPages("")
	// pdf.SetHeaderFunc(doc.headerFunc)
	// pdf.SetFooterFunc(doc.footerFunc)

	doc.document = pdf
	return doc
}

// func (r *PDFDoc) Method(content string) {
// 	fs := 10.0
// 	r.document.SetFontSize(fs)

// 	// content := "Captura de exons com Nextera Exome Capture seguida por sequenciamento de nova " +
// 	// 	"geração com Illumina HiSeq. Alinhamento e identificação de variantes utilizando protocolos " +
// 	// 	"de bioinformática, tendo como referência a versão GRCh37 do genoma humano. Análise médica " +
// 	// 	"orientada pelas informações que motivaram a realização deste exame."

// 	r.drawLine(fs, "<b>Método</b>", 10)

// 	// MultiCell(width, height, content, border, align, fill)
// 	r.document.MultiCell(0, 5, r.encodingFunc(content), "", "", false)
// 	r.lineStroke()
// }

// url := "https://cdn.recast.ai/newsletter/city-01.png"
// httpimg.Register(pdf, url, "")
// pdf.Image(url, 15, 15, 267, 0, false, "", 0, "")
