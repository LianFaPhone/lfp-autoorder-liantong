package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var GConfig Config
var GPreConfig PreConfig

func LoadConfig(path string) *Config {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("Read yml config[%s] err[%v]", path, err).Error())
	}

	err = yaml.Unmarshal([]byte(data), &GConfig)
	if err != nil {
		panic(fmt.Errorf("yml.Unmarshal config[%s] err[%v]", path, err).Error())
	}
	err = PreProcess()
	if err != nil {
		panic(fmt.Errorf("PreProcess config[%s] err[%v]", path, err).Error())
	}
	return &GConfig
}

func PreProcess() error {
	GPreConfig.ShopUrlMap = make(map[string]*ShopUrl)
	for i:=0; i < len(GConfig.ShopUrl);i++ {
		GPreConfig.ShopUrlMap[GConfig.ShopUrl[i].Num] =GConfig.ShopUrl[i]
	}
	GPreConfig.CardTypeMap = make(map[string]*CardType)
	for i:=0; i < len(GConfig.CardType);i++ {
		GPreConfig.CardTypeMap[GConfig.CardType[i].Num] =GConfig.CardType[i]
	}
	ok := false
	GPreConfig.NowShopUrl,ok  = GPreConfig.ShopUrlMap[GConfig.Server.ShopId]
	if !ok {
		return fmt.Errorf("shop_id config err")
	}
	return nil
}

type PreConfig struct{
	ShopUrlMap map[string] *ShopUrl
	CardTypeMap map[string] *CardType
	NowShopUrl *ShopUrl
}

type Config struct {
	Server  System      `yaml:"system"`
	ShopUrl []*ShopUrl     `yaml:"shop_urls"`
	CardType []*CardType   `yaml:"card_type"`
}

type System struct {
	OutputPath    string    `yaml:"output_path"`
	InputPath   string        `yaml:"input_path"`
	LogPath string      `yaml:"log_path"`
	PicPath string      `yaml:"pic_path"`
	Intvl    int      `yaml:"interval"`
	CardId   string     `yaml:"card_id"`
	ShopId   string     `yaml:"shop_id"`
}

type ShopUrl struct {
	Name       string  `yaml:"name"`
	Url        string  `yaml:"url"`
	Num        string  `yaml:"id"`
	Host        string  `yaml:"host"`
	SchemeData        string  `yaml:"schemeData"`
	Phone        string  `yaml:"phone"`
}

type CardType struct{
	Name       string  `yaml:"name"`
	Num        string  `yaml:"id"`
	OrderTp        string  `yaml:"orderType"`
}

