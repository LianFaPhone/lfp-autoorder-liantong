package main

import(
	"LianFaPhone/lfp-common"
	"fmt"
	"net/url"
	"encoding/json"
	. "LianFaPhone/lfp-autoorder-liantong/config"
)


type (
	ReGzt struct{

	}
	ResGzt struct{
		Total   int   `json:"total"`
		Success  bool   `json:"success"`
		Message string `json:"message"`
	}
)

func (this * ReGzt) Send(name, idcard string) (bool, string,error) {
	shopUrl,ok  := GPreConfig.ShopUrlMap[GConfig.Server.ShopId]
	if !ok {
		return false,"", fmt.Errorf("shop_id config err")
	}
	heads := map[string]string{
		"Referer": shopUrl.Url,
	}

	formBody := make(url.Values)
	formBody.Add("certName", name)
	formBody.Add("certNum", idcard)


	resData,err := common.HttpFormSend(shopUrl.Host+"/broadBandInstall/checkGZT.do", formBody, "POST", heads)
	if err != nil {
		return false,"",err
	}
	res:=new(ResGzt)
	if err = json.Unmarshal(resData, res); err != nil {
		return false,"",err
	}
	if !res.Success {
		return false,fmt.Sprintf("%v-%v", res.Success, res.Message),nil
	}
	return true,"", nil
}