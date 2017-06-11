package engine

import (
	"errors"
	"fmt"
	"time"
)

// SimpanData ..
func SimpanData(datanya DataSerial) error {

	var temp = DataSerialDB{
		Busvoltage:   datanya.Busvoltage,
		Current:      datanya.Current,
		Humidity:     datanya.Humidity,
		Loadvoltage:  datanya.Loadvoltage,
		Shuntvoltage: datanya.Shuntvoltage,
		Temperature:  datanya.Temperature,
		UpdateAt:     time.Now().String(),
	}
	// db.Create(&animal)
	errr := KonDB.Create(&temp).GetErrors()
	if len(errr) > 0 {
		fmt.Println(errr)
		var stringErr string
		for index := 0; index < len(errr); index++ {
			stringErr = stringErr + "," + string(errr[index].Error())
		}
		return errors.New(stringErr)
	}
	return nil
}
