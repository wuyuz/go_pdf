package services

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"path/filepath"
	"pdfUtil/lib"
)

var inDir  = "../data"

type PdfService struct {
	msg             string
	inFile, outFile string
	selectedPages   []string
	mode            string
	modeParm        string
	wmConf          string
}


func NewPdfService() *PdfService{
	return &PdfService{
		msg:           "PDF merge function",
		inFile:        "",
		outFile:       "",
		selectedPages: nil,
		mode:          "pdf",
		modeParm:      "",
		wmConf:        "sc:1, pos:tr, off:-10 0, rot:0",
	}
}

func (p *PdfService) AlertAttr(in,out,modeParm string){
	p.inFile = in
	p.outFile = out
	p.modeParm = modeParm
}

func (p *PdfService) AddWatermarks(onTop bool) error {
	var err error
	switch p.mode {
	case "text":
		err = api.AddTextWatermarksFile(p.inFile, p.outFile, p.selectedPages, onTop, p.modeParm,p.wmConf, nil)
	case "image":
		err = api.AddImageWatermarksFile(p.inFile, p.outFile, p.selectedPages, onTop, p.modeParm, p.wmConf, nil)
	case "pdf":
		err = api.AddPDFWatermarksFile(p.inFile, p.outFile, p.selectedPages, onTop, p.modeParm, p.wmConf, nil)
	}
	if err != nil {
		return err
	}
	if err := api.ValidateFile(p.outFile, nil); err != nil {
		return err
	}
	return nil
}


func AddWatermarks(msg, inFile, outFile string, selectedPages []string, mode, modeParam, desc string, onTop bool) {
	inFile = filepath.Join(inDir, inFile)

	s := "watermark"
	if onTop {
		s = "stamp"
	}
	outFile = filepath.Join("../data/out", s, mode, outFile)
	fmt.Println(inFile,outFile)
	var err error
	switch mode {
	case "text":
		err = api.AddTextWatermarksFile(inFile, outFile, selectedPages, onTop, modeParam, desc, nil)
	case "image":
		err = api.AddImageWatermarksFile(inFile, outFile, selectedPages, onTop, modeParam, desc, nil)
	case "pdf":
		err = api.AddPDFWatermarksFile(inFile, outFile, selectedPages, onTop, modeParam, desc, nil)
	}
	if err != nil {
		lib.CheckErr(err, msg)
	}
	if err := api.ValidateFile(outFile, nil); err != nil {
		lib.CheckErr(err, msg)
	}
}


