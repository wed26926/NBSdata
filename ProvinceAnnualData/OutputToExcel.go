package ProvinceAnnualData

import (
	"github.com/tealeg/xlsx"
	"log"
)

func OutputToExcel(response Response){
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
	xlsxfile.Save("test.xlsx")
}