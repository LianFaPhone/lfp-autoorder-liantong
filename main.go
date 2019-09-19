package main


import (
	"go.uber.org/zap"
	"runtime/debug"
	"fmt"
	"LianFaPhone/lfp-base/config"
	. "LianFaPhone/lfp-base/log/zap"
	. "LianFaPhone/lfp-autoorder-liantong/config"
	"time"
	"os/exec"
	"github.com/tealeg/xlsx"
	"strings"
)
//
//func main(){
//	//读取命令行
//	//读取配置
//
//	//读取excel
//	//调用接口
//	//输出excel
//
//
//}

var AllExcel    *Excel
var SecondExcel *Excel
var SuccExcel   *Excel
var FailExcel   *Excel
var timeNowStr  string

func PanicPrint() {
	if err := recover(); err != nil {
		ZapLog().With(zap.Any("error", err)).Error(string(debug.Stack()))
	}
}

func main() {
	laxFlag := config.NewLaxFlagDefault()
	cfgPath := laxFlag.String("conf_path", "config.yaml", "config path")
	logPath := laxFlag.String("log_path", "zap.conf", "log conf path")
	laxFlag.LaxParseDefault()
	fmt.Printf("command param: conf_path=%s, log_path=%s\n", *cfgPath, *logPath)
	LoadConfig(*cfgPath)
	LoadZapConfig(*logPath)
	ZapLog().Sugar().Infof("Config Content[%v]", GConfig)
	defer ZapClose()
	defer PanicPrint()

	ZapLog().Sugar().Infof("******************Config shopNname[%s] cardTp[%s], and sleep 5 second to continue", GPreConfig.NowShopUrl.Name, GConfig.Server.CardId)
	time.Sleep(time.Second * 5)
	timeNowStr = time.Now().Format("2006-01-02-15-04-05")
	if err := Init(); err != nil {
		//老文件文件名改成日期
		ZapLog().Sugar().Errorf("Init err %v,  so exist",err)
		return
	}


	ZapLog().Sugar().Info("*****************Start Read input excel........"+timeNowStr)
	colnames, sheet, err := ReadInput()
	if err != nil {
		ZapLog().Sugar().Errorf("ReadInput err %v,  so exist",err)
		return
	}
	outcolnames := make([]string, 0)
	outcolnames = append(outcolnames, colnames...)
	outcolnames = append(outcolnames, "套餐", "状态", "错误信息")

	if err = AllExcel.AddHeader(outcolnames); err != nil {
		ZapLog().Sugar().Errorf("AllExcel.AddHeader err %v,  so exist",err)
		return
	}
	if err = SecondExcel.AddHeader(outcolnames); err != nil {
		ZapLog().Sugar().Errorf("SecondExcel.AddHeader err %v,  so exist",err)
		return
	}
	if err = SuccExcel.AddHeader(outcolnames); err != nil {
		ZapLog().Sugar().Errorf("SuccExcel.AddHeader err %v,  so exist",err)
		return
	}
	if err = FailExcel.AddHeader(outcolnames); err != nil {
		ZapLog().Sugar().Errorf("FailExcel.AddHeader err %v,  so exist",err)
		return
	}
	//if err = new(InputIndex).Parse(outcolnames); err != nil {
	//	ZapLog().Sugar().Errorf("InputIndex).Parse err %v,  so exist",err)
	//	return
	//}
	ZapLog().Sugar().Infof("**Start Read excel record [[%d]], and sleep 5 second to continue", len(sheet.Rows)-1)
	time.Sleep(time.Second * 5)

	//执行操作

	if err := Work(outcolnames, sheet); err != nil {
		ZapLog().Sugar().Errorf("work err %v so exist", err)
		return
	}
	//输出文件

	ZapLog().Sugar().Info("\n\n**************Success done*****************")
}

func Init() error {
	cmd := exec.Command("mkdir", "-p", GConfig.Server.OutputPath)
	err := cmd.Run()
	if err != nil {
		ZapLog().Sugar().Error("mkdir %v err %v", GConfig.Server.OutputPath, err)
		return err
	}

	//cmd = exec.Command("mkdir", "-p", GConfig.Server.InputPath)
	//err = cmd.Run()
	//if err != nil {
	//	ZapLog().Sugar().Error("mkdir %v err %v", GConfig.Server.InputPath, err)
	//	return err
	//}

	cmd = exec.Command("mkdir", "-p", GConfig.Server.PicPath)
	err = cmd.Run()
	if err != nil {
		ZapLog().Sugar().Error("mkdir %v err %v", GConfig.Server.PicPath, err)
		return err
	}
	AllExcel, err = NewExcel("", GConfig.Server.OutputPath + "/all-"+timeNowStr+".xlsx")
	if err != nil {
		ZapLog().Sugar().Error("NewExcel %v err %v", GConfig.Server.OutputPath + "/all-"+timeNowStr+".xlsx", err)
		return err
	}

	SecondExcel, err = NewExcel("", GConfig.Server.OutputPath + "/second-"+timeNowStr+".xlsx")
	if err != nil {
		ZapLog().Sugar().Error("NewExcel %v err %v", GConfig.Server.OutputPath + "/second-"+timeNowStr+".xlsx", err)
		return err
	}

	SuccExcel, err = NewExcel("", GConfig.Server.OutputPath + "/success-"+timeNowStr+".xlsx")
	if err != nil {
		ZapLog().Sugar().Error("NewExcel %v err %v", GConfig.Server.OutputPath + "/success-"+timeNowStr+".xlsx", err)
		return err
	}
	FailExcel, err = NewExcel("", GConfig.Server.OutputPath + "/fail-"+timeNowStr+".xlsx")
	if err != nil {
		ZapLog().Sugar().Error("NewExcel %v err %v", GConfig.Server.OutputPath + "/fail-"+timeNowStr+".xlsx", err)
		return err
	}

	return nil
}

func ReadInput() ([]string, *xlsx.Sheet, error) {
	excelFileName := GConfig.Server.InputPath
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		return nil,nil,err
	}
	sheet, ok := xlFile.Sheet["Sheet1"]
	if !ok {
		return nil,nil, fmt.Errorf("nofind Sheet1")
	}

	rowNames := make([]string, 0)
	for _, cell := range sheet.Rows[0].Cells {
		text := cell.String()
		text = strings.Replace(text, " ", "", -1)
		rowNames = append(rowNames, text)
	}
	if len(rowNames) <= 1 {
		return nil,nil, fmt.Errorf("in.xlsx format err")
	}
	return rowNames, sheet, nil
}


func Work(colnames []string, sheet *xlsx.Sheet) error {
	//读取
	//创建输出文件 总表，成功，失败， 二次

	orderTpCode,err := new(ReOrderType).Get()
	if err != nil {
		ZapLog().Sugar().Error("ReOrderType.Get error %v, so exist", err)
		return err
	}

	CardType, ok := GPreConfig.CardTypeMap[GConfig.Server.CardId]
	if !ok {
		ZapLog().Sugar().Error("CardTypeMap nofind error, so exist")
		return fmt.Errorf("CardTypeMap nofind error")
	}


	ZapLog().Sugar().Infof("***********orderTp [[%v]], cardName [[%v]]", orderTpCode, CardType.Name)
	fmt.Println("\n============================================\n")
	time.Sleep(time.Second *3)
	AllRecords := len(sheet.Rows) - 1
	SuccessRecords := 0
	FailRecords := 0

	for i:=0; i< len(sheet.Rows); i++ {
		if i == 0 { //第一行标题栏
			continue
		}
		row := sheet.Rows[i]

		inIdCard := row.Cells[2].String()
		inXinMing := row.Cells[1].String()
		inProName := row.Cells[4].String()
		inCityName := row.Cells[5].String()
		inQuName := row.Cells[6].String()
		inPhone := row.Cells[3].String()
		inAddress := row.Cells[8].String()
		//inStreet := row.Cells[7].String()
		quIndex :=  strings.Index(inAddress, inQuName)
		if quIndex >= 0 {
			inAddress = inAddress[quIndex+len(inQuName):]
		}

		excelArr := CellsToArr(row.Cells, CardType.Name)
		certAddr := inCityName+inQuName+inAddress
		fmt.Println("\n============================================")
		ZapLog().Sugar().Infof("************record= [%v] [%v]", excelArr, certAddr)


		if ok := VerifyXingMing(inXinMing, inIdCard, excelArr); !ok {
			FailRecords++
			continue
		}
		ZapLog().Sugar().Infof("VerifyXingMing Ok")

		proCode,cityCode,quCode,ok := GetAreaCode(inProName,inCityName,inQuName, excelArr)
		if !ok {
			FailRecords++
			continue
		}

		ZapLog().Sugar().Infof("area code %v %v %v  %v %v %v", inProName,inCityName,inQuName, proCode,cityCode,quCode)

		//----------------------------
		cardSearchArr,err := new(ReCardSearch).Send(orderTpCode)
		if err !=  nil {
			ZapLog().Sugar().Errorf("ReCardSearch.Send %v err=%v",orderTpCode, err)
			excelArr = append(excelArr, "网络问题", err.Error())
			if err2 := AllExcel.Append(excelArr); err2 != nil {
				ZapLog().Sugar().Error("AllExcel.Append error %v", err2)
			}
			if err2 := FailExcel.Append(excelArr); err2 != nil {
				ZapLog().Sugar().Error("FailExcel.Append error %v", err2)
			}
			FailRecords++
			continue
		}
		oneFlag:= false
		zanPhone := ""
		for j:=0; j< len(cardSearchArr); j++ {
			ok, err := new(ReCloseNumber).Send(cardSearchArr[j].NumId, inIdCard)
			if err != nil {
				ZapLog().Sugar().Errorf("ReCloseNumber %v %v err:%v",cardSearchArr[j].NumId,inIdCard, err)
				continue
			}
			if ok {
				oneFlag= true
				zanPhone = cardSearchArr[j].NumId
				ZapLog().Sugar().Infof("ReCloseNumber ok, %v", cardSearchArr[j].NumId)
				break
			}
		}
		if !oneFlag {
			ZapLog().Sugar().Errorf("ReCloseNumber fail")
			excelArr = append(excelArr, "占号失败")
			if err2 := AllExcel.Append(excelArr); err2 != nil {
				ZapLog().Sugar().Error("AllExcel.Append error %v", err2)
			}
			if err2 := FailExcel.Append(excelArr); err2 != nil {
				ZapLog().Sugar().Error("FailExcel.Append error %v", err2)
			}
			FailRecords++
			continue
		}


		orderId, err := new(ReOrderSubmit).Send(inXinMing, inIdCard, inPhone, inAddress, zanPhone, proCode, cityCode, quCode, certAddr, GPreConfig.NowShopUrl.Phone, orderTpCode)
		if err != nil{
			ZapLog().Sugar().Errorf("ReOrderSubmit Fail %v", err)
			secondFlag, succFlag, errMsg := OrderSubmitErr(err.Error())

			excelArr = append(excelArr, errMsg, err.Error())
			if err2 := AllExcel.Append(excelArr); err2 != nil {
				ZapLog().Sugar().Error("AllExcel.Append error %v", err2)
			}
			if err2 := FailExcel.Append(excelArr); err2 != nil {
				ZapLog().Sugar().Error("FailExcel.Append error %v", err2)
			}
			if succFlag{
				if err2 := SuccExcel.Append(excelArr); err2 != nil {
					ZapLog().Sugar().Error("SuccExcel.Append error %v", err2)
				}
			}
			if secondFlag {
				if err2 := SecondExcel.Append(excelArr); err2 != nil {
					ZapLog().Sugar().Error("SecondExcel.Append error %v", err2)
				}
			}
			FailRecords++
			continue
		}else{
			excelArr = append(excelArr, zanPhone, "成功")
			if err2 := AllExcel.Append(excelArr); err2 != nil {
				ZapLog().Sugar().Error("AllExcel.Append error %v", err2)
			}
			if err2 := SuccExcel.Append(excelArr); err2 != nil {
				ZapLog().Sugar().Error("SuccExcel.Append error %v", err2)
			}
			ZapLog().Sugar().Infof("**********ReOrderSubmit Success,order[%s]newPhone [%v]",orderId, zanPhone)
		}
		SuccessRecords++
		if err := new(RePicSubmit).Send(orderId); err != nil {
			ZapLog().Sugar().Errorf("pic submit orderId[%v] err %v", orderId, err)
		}

		if GConfig.Server.Intvl > 0 {
			time.Sleep(time.Millisecond * time.Duration(GConfig.Server.Intvl))
		}
	}
	ZapLog().Sugar().Infof("****The End**********AllRecords[%v] SuccessRecords[%v] FailRecords[%v] ", AllRecords, SuccessRecords, FailRecords)
	return nil
}

//secondFlag, succFlag
func OrderSubmitErr(err string) (bool, bool, string) {
	if strings.Contains(err, "已预约过") {
		return false,false,"已预约"
	}
	if strings.Contains(err, "已超过五户") {
		return false,false,"一证五户"
	}
	if strings.Contains(err, "专属卡超限") {
		return false,false,"专属卡超限"
	}
	if strings.Contains(err, "二次办理") {
		return true,false,"二次办理"
	}

	return false, false, err
}

//检测姓名身份证，返回 是否成功
func VerifyXingMing(inXinMing, inIdCard string, excelArr []string) bool {
	ok,msg,err :=  new(ReGzt).Send(inXinMing, inIdCard)
	if err != nil {
		ZapLog().Sugar().Errorf("ReOrderType.Get error %v %v", err, msg)
		excelArr = append(excelArr, "网络问题", msg)
		if err2 := AllExcel.Append(excelArr); err2 != nil {
			ZapLog().Sugar().Error("AllExcel.Append error %v", err)
		}
		if err2 := FailExcel.Append(excelArr); err2 != nil {
			ZapLog().Sugar().Error("FailExcel.Append error %v", err2)
		}
		return false
	}
	if !ok {//身份校验失败
		ZapLog().Sugar().Errorf("name and idcard not verified, %v", msg)
		excelArr = append(excelArr, "身份证校验失败", msg)
		if err2 := AllExcel.Append(excelArr); err2 != nil {
			ZapLog().Sugar().Error("AllExcel.Append error %v", err)
		}
		if err2 := FailExcel.Append(excelArr); err2 != nil {
			ZapLog().Sugar().Error("FailExcel.Append error %v", err2)
		}
		return false
	}
	return true
}

func GetAreaCode(inProName,inCityName,inQuName string, excelArr []string) (string,string,string,bool) {
	provice, ok := ProviceMap[inProName]
	if !ok {
		excelArr = append(excelArr, "地址未找到", "省")
		ZapLog().Sugar().Errorf("nofind provice")
		if err2 := AllExcel.Append(excelArr); err2 != nil {
			ZapLog().Sugar().Error("AllExcel.Append error %v", err2)
		}
		if err2 := FailExcel.Append(excelArr); err2 != nil {
			ZapLog().Sugar().Error("FailExcel.Append error %v", err2)
		}
		return "","","",false
	}
	proCode := provice.Code
	mm,err := new(ReqArea).Get("MALL_SHIP_CLOUD_AREA", proCode) //省获取市列表
	if err != nil {
		ZapLog().Sugar().Errorf("ReqArea.Get MALL_SHIP_CLOUD_AREA %v err=%v",proCode, err)
		excelArr = append(excelArr, "网络问题", err.Error())
		if err2 := AllExcel.Append(excelArr); err2 != nil {
			ZapLog().Sugar().Error("AllExcel.Append error %v", err2)
		}
		if err2 := FailExcel.Append(excelArr); err2 != nil {
			ZapLog().Sugar().Error("FailExcel.Append error %v", err2)
		}
		return "","","",false
	}

	cityArea,ok := mm[inCityName]
	if !ok {
		ZapLog().Sugar().Errorf("city nofind %v",inCityName)
		excelArr = append(excelArr, "地址未找到", "市")
		if err2 := AllExcel.Append(excelArr); err2 != nil {
			ZapLog().Sugar().Error("AllExcel.Append error %v", err2)
		}
		if err2 := FailExcel.Append(excelArr); err2 != nil {
			ZapLog().Sugar().Error("FailExcel.Append error %v", err2)
		}
		return "","","",false
	}
	cityCode := cityArea.PValue
	mm,err = new(ReqArea).Get("MALL_SHIP_CLOUD_COUNTY", cityArea.PValue)
	if err != nil {
		ZapLog().Sugar().Errorf("ReqArea.Get MALL_SHIP_CLOUD_COUNTY %v err=%v",cityCode, err)
		excelArr = append(excelArr, "网络问题", err.Error())
		if err2 := AllExcel.Append(excelArr); err2 != nil {
			ZapLog().Sugar().Error("AllExcel.Append error %v", err2)
		}
		if err2 := FailExcel.Append(excelArr); err2 != nil {
			ZapLog().Sugar().Error("FailExcel.Append error %v", err2)
		}
		return "","","",false
	}

	quArea, ok := mm[inQuName]
	if !ok {
		ZapLog().Sugar().Errorf("qu nofind %v",inQuName)
		excelArr = append(excelArr, "地址未找到", "区")
		if err2 := AllExcel.Append(excelArr); err2 != nil {
			ZapLog().Sugar().Error("AllExcel.Append error %v", err2)
		}
		if err2 := FailExcel.Append(excelArr); err2 != nil {
			ZapLog().Sugar().Error("FailExcel.Append error %v", err2)
		}
		return "","","",false
	}
	quCode := quArea.PValue

	return proCode,cityCode,quCode,true
}