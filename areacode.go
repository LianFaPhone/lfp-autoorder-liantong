package main

import(
	"LianFaPhone/lfp-common"
	. "LianFaPhone/lfp-autoorder-liantong/config"
	"fmt"
	"encoding/json"
)

var areaCodeMap = make(map[string] map[string]*AreaCode)


type (
	ReqArea struct{

	}

	ResArea struct{
		Total   int   `json:"total"`
		Success  bool   `json:"success"`
		Message string `json:"message"`
		Data []*AreaCode `json:"data"`
		/*
		pDesc: "上海市"
pKey: "MALL_SHIP_CLOUD_AREA"
pPValue: "310000"
pValue: "310100"
		 */
	}
	AreaCode struct{
		PDesc   string `json:"pDesc"`  //中文名称
		PKey    string `json:"pKey"`  //"MALL_SHIP_CLOUD_AREA"
		PPValue string `json:"pPValue"` //父亲
		PValue  string `json:"pValue"`
	}
)

func (this *ReqArea) Get(key, parent_code string) (map[string]*AreaCode, error){
	arr,ok := areaCodeMap[parent_code]
	if ok {
		return arr,nil
	}
	res,err := this.Send(key, parent_code)
	if err != nil {
		return nil,err
	}
	if len(res) == 0 {
		return nil,fmt.Errorf("nil resp")
	}
	mm := make(map[string]*AreaCode)
	for i:=0; i< len(res);i++{
		mm[res[i].PDesc] = res[i]
	}
	areaCodeMap[parent_code] = mm
	return mm,nil

}

func (this *ReqArea) Send(key, parent_code string) ([]*AreaCode, error){
	shopUrl,ok  := GPreConfig.ShopUrlMap[GConfig.Server.ShopId]
	if !ok {
		return nil,fmt.Errorf("shop_id config err")
	}
	heads := map[string]string{
		"Referer": shopUrl.Url,
	}
	data,err := common.HttpSend(shopUrl.Host+"/common/getCityOrCounty.do?p_key="+key+"&p_value="+parent_code, nil, "POST", heads)
	if err != nil {
		return nil,err
	}
	res := new(ResArea)

	err = json.Unmarshal(data, res)
	if err != nil {
		return nil,err
	}
	if !res.Success {
		return nil,fmt.Errorf("%v-%v", res.Success, res.Message)
	}
	return res.Data,nil
}
