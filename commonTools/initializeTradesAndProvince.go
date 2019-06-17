package commonTools

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type TradeTree struct{
	Name   string      `json:"name"`
	Id string          `json:"id"`
	Childs []TradeTree `json:"childs"`
}

var province map[string]string
var root TradeTree

func init() {
	initializeProvince()
	initializeTrade()
}

func initializeProvince(){
	province = make(map[string]string)
	content,err := ioutil.ReadFile(os.Getenv("GOPATH")+"/src/NBSdata/commonTools/province.txt")
	if err != nil{
		panic(err)
	}
	provincedata := strings.Split(string(content),"\n")
	for _,data := range provincedata{
		provinceDetail := strings.Split(data," ")
		province[provinceDetail[0]] = provinceDetail[1]
	}
}

func initializeTrade(){
	content,err := ioutil.ReadFile(os.Getenv("GOPATH")+"/src/NBSdata/commonTools/trades.json")
	if err != nil{
		panic(err)
	}
	err = json.Unmarshal(content,&root)
	if err != nil{
		panic(err)
	}
}
func  GetProvinceData() map[string]string{
	return province
}

func GetTrade() TradeTree {
	return root
}