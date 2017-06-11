package models

import (
	"fmt"
	"log"

	"github.com/edot92/simple_datalogger_baterai/engine"
)

type ParamChart struct {
	JenisChart string `json:"jenis_chart"`
}
type ResponChart struct {
	Value float32 `json:"value"`
	Time  string  `json:"time"`
}
type Paramhistorydate struct {
	Startdate string `json:"startdate"`
	Enddate   string `json:"enddate"`
}

func GetRealtimeChart(jenisChart string) ([]ResponChart, error) {
	var res []engine.DataSerialDB
	rows, err := engine.KonDB.DB().Query("select  count(*) FROM  data_serial_dbs")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	rows.Next()
	var count int
	err = rows.Scan(&count)
	if err != nil {
		panic(err)
	}
	fmt.Printf("count: %d\n", count)
	engine.KonDB.Limit(10).Offset(count - 10).Find(&res)
	var resChartnya []ResponChart
	for index := 0; index < len(res); index++ {
		resChartnya = append(resChartnya, ResponChart{
			res[index].Temperature,
			res[index].CreatedAt,
		})
	}
	fmt.Println(resChartnya)
	return resChartnya, nil
}
func GetRealtimeBox() (engine.DataSerialDB, error) {
	var res engine.DataSerialDB
	engine.KonDB.Last(&res)
	return res, nil

}

func GetHistoryLast() ([]engine.DataSerialDB, error) {
	rows, err := engine.KonDB.DB().Query("select  count(*) FROM  data_serial_dbs")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	rows.Next()
	var count int
	err = rows.Scan(&count)
	if err != nil {
		panic(err)
	}
	fmt.Printf("count: %d\n", count)
	var res []engine.DataSerialDB
	engine.KonDB.Limit(10).Offset(count - 10).Find(&res)
	// var resChartnya []ResponChart
	// for index := 0; index < len(res); index++ {
	// 	resChartnya = append(resChartnya, ResponChart{
	// 		res[index].Temperature,
	// 		res[index].CreatedAt.Unix(),
	// 	})
	// }
	// fmt.Println(resChartnya)
	return res, nil
}

func GetHistoryDateRange(param Paramhistorydate) ([]engine.DataSerialDB, error) {
	var res []engine.DataSerialDB
	// param.Startdate = "2016-01-24T12:00:00.000000000+07:00"
	// param.Enddate = "2019-01-24T12:00:00.000000000+07:00"
	// err := engine.KonDB.Debug().Where("created_at >= ? AND created_at <=?", param.Startdate, param.Enddate).Find(&res).Error
	err := engine.KonDB.Debug().Where(" created_at BETWEEN  (?) AND (?) ", param.Startdate, param.Enddate).Find(&res).Error

	if err != nil {
		log.Fatal(err)
	}
	return res, nil

}
