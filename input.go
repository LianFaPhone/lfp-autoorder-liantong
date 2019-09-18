package main

import "fmt"

type InputIndex struct{
	ProI  int
	CityI int
	QuI   int
	StreetI int
	AddrI int
	XimingI int
	PhoneI  int
	IdcardI int
}

func (this *InputIndex) Parse(colnames []string)  error {
	this.ProI = GetIndexFromArr(colnames, "省")
	if this.ProI < 0 {
		return fmt.Errorf("in.xlsx format err")
	}
	this.CityI = GetIndexFromArr(colnames, "市")
	if this.CityI < 0 {
		return fmt.Errorf("in.xlsx format err")
	}
	this.QuI = GetIndexFromArr(colnames, "区")
	if this.QuI < 0 {
		return fmt.Errorf("in.xlsx format err")
	}
	 this.StreetI = GetIndexFromArr(colnames, "街道")
	if this.StreetI < 0 {
		return fmt.Errorf("in.xlsx format err")
	}
	  this.AddrI = GetIndexFromArr(colnames, "地址")
	if this.AddrI < 0 {
		return fmt.Errorf("in.xlsx format err")
	}
	this.XimingI = GetIndexFromArr(colnames, "姓名")
	if this.XimingI < 0 {
		return fmt.Errorf("in.xlsx format err")
	}
	this.PhoneI = GetIndexFromArr(colnames, "电话号码")
	if this.PhoneI < 0 {
		return fmt.Errorf("in.xlsx format err")
	}
	this.IdcardI = GetIndexFromArr(colnames, "身份证号码")
	if this.IdcardI < 0 {
		return fmt.Errorf("in.xlsx format err")
	}
	return nil
}
