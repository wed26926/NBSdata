package ProvinceAnnualData

import (
	"github.com/tealeg/xlsx"
	"log"
)

func OutputToExcel(response Response,filepath string){
	xlsxfile := xlsx.NewFile()
	sheet,err := xlsxfile.AddSheet("指标")
	if err != nil{
		log.Println(err)
	}
	row1 := sheet.AddRow()
	zbnames := make([]map[string]string,len(response.Returndata.Wdnodes))
	zblocations := make(map[string]int)
	for i,wd := range response.Returndata.Wdnodes{
		row1.AddCell().Value = wd.Wdname
		zblocations[wd.Wdcode] = i
		zbnames[i] = make(map[string]string)
		for _,wdnode := range wd.Nodes{
			zbnames[i][wdnode.Code] = wdnode.Name
		}
	}
	row1.AddCell().Value = "值"
	for _,data := range response.Returndata.Datanodes{
		row := sheet.AddRow()
		row.Cells = make([]*xlsx.Cell,len(zbnames))
		for _,wd := range data.Wds{
			row.Cells[zblocations[wd.Wdcode]] = new(xlsx.Cell)
			row.Cells[zblocations[wd.Wdcode]].Value = zbnames[zblocations[wd.Wdcode]][wd.Valuecode]
		}
		row.AddCell().SetFloat(data.Data.Data)
	}
	sheet2,_ := xlsxfile.AddSheet("sheet2")
	row2 := sheet2.AddRow()
	row2.AddCell().Value = "指标"
	row2.AddCell().Value = "单位"
	for _,wd := range response.Returndata.Wdnodes{
		if wd.Wdcode != "zb"{
			continue
		}
		for _,data := range wd.Nodes{
			row := sheet2.AddRow()
			row.AddCell().Value = data.Name
			row.AddCell().Value = data.Unit
		}
		row := sheet2.AddRow()
		row.AddCell()
	}
	xlsxfile.Save(filepath)
}

func MultiOut(responses []Response,Dirname string){

}