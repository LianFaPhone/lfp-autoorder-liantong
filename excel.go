package main

import "github.com/tealeg/xlsx"


func NewExcel(sheetName, filePath string) (*Excel, error){
	e := new(Excel)
	err := e.Init(sheetName, filePath)
	if err != nil {
		return nil,err
	}
	return e,nil
}

type Excel struct{
	file    *xlsx.File
	sheet   *xlsx.Sheet
	rowNames []string
	filePath string
}

func (this *Excel) Init(sheetName, filePath string) (err error) {
	this.filePath = filePath
	this.file = xlsx.NewFile()
	if len(sheetName) <= 0 {
		sheetName = "Sheet1"
	}
	this.sheet, err = this.file.AddSheet(sheetName)

	if err != nil {
		return
	}
	return
}

func (this *Excel) AddHeader(data [] string) error {
	row := this.sheet.AddRow()
	for i:=0; i < len(data); i++ {
		row.AddCell().SetValue(data[i])
	}
	if err := this.file.Save(this.filePath); err != nil {
		return err
	}
	this.rowNames = data
	return nil
}

func (this *Excel) Append(data [] string) error {
	row := this.sheet.AddRow()
	for i:=0; i < len(data); i++ {
		row.AddCell().SetValue(data[i])
	}
	if err := this.file.Save(this.filePath); err != nil {
		return err
	}
	return nil
}

func (this *Excel) AppendCells(Cells  []*xlsx.Cell) error {
	row := this.sheet.AddRow()
	for i:=0; i < len(Cells); i++ {
		row.AddCell().SetValue(Cells[i].Value)
	}
	if err := this.file.Save(this.filePath); err != nil {
		return err
	}
	return nil
}

func CellsToArr(Cells  []*xlsx.Cell, taoCan string) ([]string) {
	arr :=make([]string, 0, len(Cells))
	for i:=0; i < len(Cells); i++ {
		arr = append(arr, Cells[i].Value)
	}
	arr =append(arr, taoCan)
	return arr
}
