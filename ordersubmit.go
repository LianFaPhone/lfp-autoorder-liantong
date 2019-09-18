package main


import(
	"LianFaPhone/lfp-common"
	"fmt"
	"net/url"
	"encoding/json"
	. "LianFaPhone/lfp-autoorder-liantong/config"
	"strings"
	"LianFaPhone/lfp-base/log/zap"
)

type (
	ReOrderSubmit struct{

	}
	ResOrderSubmit struct{
		Total   int   `json:"total"`
		Success  bool   `json:"success"`
		Message string `json:"message"`
		Data    interface{} `json:"data"`
	}
)

func (this *ReOrderSubmit) Send(custName, certNum,phone,address,cardNum,province,areaCode,countyCode,certAddr,serviceNumber, orderType string ) (string, error){
	shopUrl,ok  := GPreConfig.ShopUrlMap[GConfig.Server.ShopId]
	if !ok {
		return  "", fmt.Errorf("shop_id config err")
	}
	heads := map[string]string{
		"Referer": shopUrl.Url,
	}

	formBody := make(url.Values)
	formBody.Add("productId", GConfig.Server.CardId)
	formBody.Add("custName", custName)
	formBody.Add("certNum", certNum)
	formBody.Add("phone", phone)
	formBody.Add("address", address)
	formBody.Add("cardNum", cardNum)
	formBody.Add("province", province)
	formBody.Add("areaId", "K")
	formBody.Add("areaCode", areaCode)
	formBody.Add("countyCode", countyCode)
	formBody.Add("schemeData", shopUrl.SchemeData)
	formBody.Add("certAddr", certAddr)
	formBody.Add("serviceNumber", serviceNumber)
	formBody.Add("orderType", orderType)
	formBody.Add("poolCode", "-1")


	resData,err := common.HttpFormSend(shopUrl.Host+"/nb_card/intentOrderSubmit.do", formBody, "POST", heads)
	if err != nil {
		return "",err
	}
	res:=new(ResOrderSubmit)
	if err = json.Unmarshal(resData, res); err != nil {
		return "",err
	}
	if !res.Success {
		return "",fmt.Errorf("%v-%v", res.Success, res.Message)
	}
	orderId := ""
	log.ZapLog().Sugar().Infof("intentOrderSubmit[%v]", res.Data)
	if orderId,ok := res.Data.(string);ok {
		return orderId, nil
	}
	mm,_ := res.Data.(map[string]interface{})
	for k,v := range mm {
		if strings.Contains(strings.ToLower(k),"order") {
			orderId = v.(string)
			return orderId,nil
		}
	}
	for k,v := range mm {
		if strings.Contains(strings.ToLower(k),"id") {
			orderId = v.(string)
			return orderId,nil
		}
	}
	return "",nil
}