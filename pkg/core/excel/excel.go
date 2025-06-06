package excel

import "github.com/xuri/excelize/v2"

func ReadCell(path, sheet, cell string) (string, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return "", err
	}
	return f.GetCellValue(sheet, cell)
}

func WriteCell(path, sheet, cell, value string) error {
	f := excelize.NewFile()
	f.SetCellValue(sheet, cell, value)
	return f.SaveAs(path)
}
