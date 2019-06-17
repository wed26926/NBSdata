package queryByProvince

import (
	"strconv"
	"time"
)

const(
	commonurl = "http://data.stats.gov.cn/easyquery.htm?"
)

type Data struct {
	Data		int 	`json:"data"`
	Dotcount	int 	`json:"dotcount"`
	Hasdata 	bool 	`json:"hasdata"`
	Strdata 	string 	`json:"strdata"`
}
type Wd struct{
	Valuecode 	string 	`json:"valuecode"`
	Wdcode		string 	`json:"wdcode"`
}
type Datanode struct {
	Code 		string	`json:"code"`
	Data		Data	`json:"data"`
	Wds			[]Wd 	`json:"wds"`
}
type Node struct {
	Cname		string	`json:"cname"`
	Code		string	`json:"code"`
	Dotcount	int		`json:"dotcount"`
	Exp 		string	`json:"exp"`
	Ifshowcode	bool	`json:"ifshowcode"`
	Memo 		string 	`json:"memo"`
	Name 		string 	`json:"name"`
	Nodesort 	string 	`json:"nodesort"`
	Sortcode 	string 	`json:"sortcode"`
	Tag 		string 	`json:"tag"`
	Unit 		string 	`json:"unit"`
}
type Wdnode struct {
	Nodes 		[]Node 	`json:"nodes"`
	Wdcode		string 	`json:"wdcode"`
	Wdname		string 	`json:"wdname"`
}
type Returndata struct {
	Datanodes 	[]Datanode	`json:"datanodes"`
	Wdnodes		[]Wdnode 	`json:"wdnodes"`
}
type Response struct {
	Returncode	int 		`json:"returncode"`
	Returndata	Returndata 	`json:"returndata"`
}

func QueryByProvince(provinceCode string,tradeCode string){
	timetemp := time.Now().UnixNano()
	k1 := strconv.FormatInt(timetemp,64)
}