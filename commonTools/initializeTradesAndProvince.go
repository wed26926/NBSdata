package commonTools

import (
	"io/ioutil"
	"os"
	"strings"
)

var province map[string]string

func init() {
	initializeProvince()
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

func  GetProvinceData() map[string]string{
	return province
}