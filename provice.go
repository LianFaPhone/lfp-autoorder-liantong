package main



var ProviceMap  map[string] *Provice


type Provice struct{
		Name string
		Code string
}

func init() {
	ProviceMap = make(map[string] *Provice)
	temp := &Provice{"北京市", "110000"}
	ProviceMap["北京市"] = temp
	ProviceMap["北京"] = temp

	temp = &Provice{"天津市", "120000"}
	ProviceMap["天津市"] = temp
	ProviceMap["天津"] = temp

	temp = &Provice{"河北省", "130000"}
	ProviceMap["河北省"] = temp
	ProviceMap["河北"] = temp

	temp = &Provice{"山西省", "140000"}
	ProviceMap["山西省"] = temp
	ProviceMap["山西"] = temp

	temp = &Provice{"内蒙古自治区", "150000"}
	ProviceMap["内蒙古自治区"] = temp
	ProviceMap["内蒙古"] = temp

	temp = &Provice{"辽宁省", "210000"}
	ProviceMap["辽宁省"] = temp
	ProviceMap["辽宁"] = temp

	temp = &Provice{"吉林省", "220000"}
	ProviceMap["吉林省"] = temp
	ProviceMap["吉林"] = temp

	temp = &Provice{"黑龙江省", "230000"}
	ProviceMap["黑龙江省"] = temp
	ProviceMap["黑龙江"] = temp

	temp = &Provice{"上海市", "310000"}
	ProviceMap["上海市"] = temp
	ProviceMap["上海"] = temp

	temp = &Provice{"江苏省", "320000"}
	ProviceMap["江苏省"] = temp
	ProviceMap["江苏"] = temp

	temp = &Provice{"浙江省", "330000"}
	ProviceMap["浙江省"] = temp
	ProviceMap["浙江"] = temp

	temp = &Provice{"安徽省", "340000"}
	ProviceMap["安徽省"] = temp
	ProviceMap["安徽"] = temp

	temp = &Provice{"福建省", "350000"}
	ProviceMap["福建省"] = temp
	ProviceMap["福建"] = temp

	temp = &Provice{"江西省", "360000"}
	ProviceMap["江西省"] = temp
	ProviceMap["江西"] = temp

	temp = &Provice{"山东省", "370000"}
	ProviceMap["山东省"] = temp
	ProviceMap["山东"] = temp

	temp = &Provice{"河南省", "410000"}
	ProviceMap["河南省"] = temp
	ProviceMap["河南"] = temp

	temp = &Provice{"湖北省", "420000"}
	ProviceMap["湖北省"] = temp
	ProviceMap["湖北"] = temp

	temp = &Provice{"湖南省", "430000"}
	ProviceMap["湖南省"] = temp
	ProviceMap["湖南"] = temp

	temp = &Provice{"广东省", "440000"}
	ProviceMap["广东省"] = temp
	ProviceMap["广东"] = temp

	temp = &Provice{"广西壮族自治区", "450000"}
	ProviceMap["广西壮族自治区"] = temp
	ProviceMap["广西"] = temp

	temp = &Provice{"海南省", "460000"}
	ProviceMap["海南省"] = temp
	ProviceMap["海南"] = temp

	temp = &Provice{"重庆市", "500000"}
	ProviceMap["重庆市"] = temp
	ProviceMap["重庆"] = temp

	temp = &Provice{"四川省", "510000"}
	ProviceMap["四川省"] = temp
	ProviceMap["四川"] = temp

	temp = &Provice{"贵州省", "520000"}
	ProviceMap["贵州省"] = temp
	ProviceMap["贵州"] = temp

	temp = &Provice{"云南省", "530000"}
	ProviceMap["云南省"] = temp
	ProviceMap["云南"] = temp

	temp = &Provice{"西藏自治区", "540000"}
	ProviceMap["西藏自治区"] = temp
	ProviceMap["西藏"] = temp

	temp = &Provice{"陕西省", "610000"}
	ProviceMap["陕西省"] = temp
	ProviceMap["陕西"] = temp

	temp = &Provice{"甘肃省", "620000"}
	ProviceMap["甘肃省"] = temp
	ProviceMap["甘肃"] = temp

	temp = &Provice{"青海省", "630000"}
	ProviceMap["青海省"] = temp
	ProviceMap["青海"] = temp

	temp = &Provice{"宁夏回族自治区", "640000"}
	ProviceMap["宁夏回族自治区"] = temp
	ProviceMap["宁夏"] = temp

	temp = &Provice{"新疆维吾尔自治区", "650000"}
	ProviceMap["新疆维吾尔自治区"] = temp
	ProviceMap["新疆"] = temp
}

