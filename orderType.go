package main


import(
	"LianFaPhone/lfp-common"
	"fmt"
	"encoding/json"
	. "LianFaPhone/lfp-autoorder-liantong/config"
)

var OrderTypeMap = make(map[string]string)
type (
	ReOrderType struct{

	}

	ResOrderType struct{
		Total   int   `json:"total"`
		Success  bool   `json:"success"`
		Message string `json:"message"`
		Data    *OrderTp   `json:"data"`
	}
	OrderTp struct{
		OrderType string  `json:"orderType"`
		PoolCode   string  `json:"poolCode"`
	}
)

func (this * ReOrderType) Get() (string,error) {
	orderTpCode,ok := OrderTypeMap[GConfig.Server.CardId]
	if ok {
		return orderTpCode,nil
	}
	orderTpCode,err := this.Send()
	if err != nil {
		return "",err
	}
	OrderTypeMap[GConfig.Server.CardId] = orderTpCode
	return orderTpCode,nil
}

func (this * ReOrderType) Send() (string, error) {
	shopUrl,ok  := GPreConfig.ShopUrlMap[GConfig.Server.ShopId]
	if !ok {
		return "",fmt.Errorf("shop_id config err")
	}
	heads := map[string]string{
		"Referer": shopUrl.Url,
	}

	resData,err := common.HttpSend(shopUrl.Host+"/nb_card/getOrderType.do?schemeData="+shopUrl.SchemeData+"&goodsId="+GConfig.Server.CardId+"&areaId=K", nil, "POST", heads)
	if err != nil {
		return "",err
	}
	res:=new(ResOrderType)
	if err = json.Unmarshal(resData, res); err != nil {
		return "",err
	}
	if !res.Success {
		return "",fmt.Errorf("%v-%v", res.Success, res.Message)
	}
	if res.Data == nil  {
		return "",nil
	}
	return res.Data.OrderType,nil
}
