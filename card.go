package main

import(
	"LianFaPhone/lfp-common"
	"fmt"
	"net/url"
	"encoding/json"
	. "LianFaPhone/lfp-autoorder-liantong/config"
)

type (
	ReCardSearch struct{

	}

	ResCardSearch struct{
		Total   int   `json:"total"`
		Success  bool   `json:"success"`
		Message string `json:"message"`
		Data []*CardSearch `json:"data"`
	}

	CardSearch struct{
		AdvancePay string     `json:"advancePay"`
		ClassId    string     `json:"classId"`
		LowCostPro  string    `json:"lowCostPro"`
		NumId string          `json:"numId"`
		//numMemo: null
		TimeDurPro  string    `json:"timeDurPro"`
	}

	ReCloseNumber struct{

	}

	ResCloseNumber struct{
		Total   int   `json:"total"`
		Success  bool   `json:"success"`
		Message string `json:"message"`
		//Data []*CardSearch `json:"data"`
	}
)

func (this * ReCardSearch) Send(orderTpCode string) ([]*CardSearch, error) {
	shopUrl,ok  := GPreConfig.ShopUrlMap[GConfig.Server.ShopId]
	if !ok {
		return nil,fmt.Errorf("shop_id config err")
	}
	heads := map[string]string{
		"Referer": shopUrl.Url,
	}

	formBody := make(url.Values)
	if orderTpCode == "" {
		formBody.Add("areaId", "K")
		formBody.Add("poolCode", "-1")
	}else{
		formBody.Add("areaId", "K")
		formBody.Add("schemeData", shopUrl.SchemeData)
		formBody.Add("orderType", orderTpCode)
	}



	resData,err := common.HttpFormSend(shopUrl.Host+"/SelfSpreadTencertCard/searchCard.do", formBody, "POST", heads)
	if err != nil {
		return nil,err
	}
	res:=new(ResCardSearch)
	if err = json.Unmarshal(resData, res); err != nil {
		return nil,err
	}
	if !res.Success {
		return nil,fmt.Errorf("%v-%v", res.Success, res.Message)
	}
	if len(res.Data) == 0 {
		return nil,fmt.Errorf("nil resp")
	}
	return res.Data,nil
}


func (this * ReCloseNumber) Send(Number, idCard string) (bool, error) {
	shopUrl,ok  := GPreConfig.ShopUrlMap[GConfig.Server.ShopId]
	if !ok {
		return false,fmt.Errorf("shop_id config err")
	}
	heads := map[string]string{
		"Referer": shopUrl.Url,
	}

	formBody := make(url.Values)
	formBody.Add("svcNum", Number)
	formBody.Add("cityCode", "K")
	formBody.Add("certNum", idCard)


	resData,err := common.HttpFormSend(shopUrl.Host+"/SelfSpreadTencertCard/choseNumber.do", formBody, "POST", heads)
	if err != nil {
		return false,err
	}
	res:=new(ResCloseNumber)
	if err = json.Unmarshal(resData, res); err != nil {
		return false,err
	}
	if !res.Success {
		return res.Success,fmt.Errorf("%v-%v", res.Success, res.Message)
	}

	return res.Success,nil
}