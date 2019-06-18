package commonTools

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func RebuildTradeTree(){
	basicTrades := [...]string{"综合","国民经济核算","人口","就业人员和工资","固定资产投资和房地产","对外经济贸易","能源","财政","价格指数","人民生活","城市概况","资源和环境","农业","工业","建筑业","运输和邮电","社会消费品零售总额","批发和零售业","住宿和餐饮业","旅游业","金融业","教育","科技","卫生","社会服务","文化","体育","公共管理、社会保障及其他"}
	basicId := []byte{65,48}
	var newRoot TradeTree
	newRoot.Level,newRoot.Name,newRoot.Id = 0,"指标",""
	for i,trade := range basicTrades{
		var id []byte
		if i > 8{
			id = append(basicId,byte(i-9+65))
		}else{
			id = append(basicId,byte(i+49))
		}
		var newTrade TradeTree
		newTrade.Id,newTrade.Name,newTrade.Level = string(id),trade,1
		newTrade.Childs = selectTrade(newTrade.Id,1)
		newRoot.Childs = append(newRoot.Childs,newTrade)
	}
	result,err := json.Marshal(newRoot)
	if err != nil{
		log.Println("重建行业树失败")
		log.Println(err)
		return
	}
	ioutil.WriteFile("trades.json",result,0644)
}

func selectTrade(id string,level int) []TradeTree{
	time.Sleep(1*time.Second)
	url := "http://data.stats.gov.cn/easyquery.htm?dbcode=fsnd&wdcode=zb&m=getTree&id="+id
	rp,err := http.Get(url)
	if err != nil{
		panic(err)
	}
	buf := bytes.NewBuffer(make([]byte,0,512))
	_,err = buf.ReadFrom(rp.Body)
	if err != nil{
		panic(err)
	}
	if strings.Contains(string(buf.Bytes()),"<html>"){
		panic("遭遇反爬措施，查询中止，请尝试适当调整查询间隔并稍后重试。")
	}
	trades := make([]TradeTree,0)
	err = json.Unmarshal(buf.Bytes(),&trades)
	if err != nil{
		log.Println(err)
	}
	if len(trades) == 0{
		return nil
	}
	for i := range trades{
		trades[i].Level = level+1
		trades[i].Childs = selectTrade(trades[i].Id,trades[i].Level)
	}
	return trades
}