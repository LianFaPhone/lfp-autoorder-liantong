package main

import(
	"go.uber.org/zap"
	. "LianFaPhone/lfp-common"
	. "LianFaPhone/lfp-base/log/zap"
)


func createXlsFile(filePath string, rowArr []interface{}, colNames []string) error {
	xlsObj,err := NewXlsx(rowArr,colNames, nil)
	if err != nil {
		ZapLog().Error("NewXlsx err", zap.Error(err))
		return err
	}
	if err = xlsObj.Generate() ; err != nil {
		ZapLog().Error("xlsObj.Generate err", zap.Error(err))
		return err
	}
	if err = xlsObj.File(filePath); err != nil {
		ZapLog().Error("xlsObj.File err", zap.Error(err))
		return err
	}
	return nil
}



func GetIndexFromArr(colnames []string, key string) int {
	for i:=0; i< len(colnames); i++ {
		if colnames[i] == key {
			return i
		}
	}
	return -1
}