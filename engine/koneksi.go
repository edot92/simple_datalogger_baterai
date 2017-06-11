package engine

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//  KonDB ....
var KonDB *gorm.DB

type DataSerialDB struct {
	ID           uint64  `gorm:"primary_key"`
	Temperature  float32 `json:"temperature"`
	Humidity     float32 `json:"humidity"`
	Current      float32 `json:"current"`
	Shuntvoltage float32 `json:"shuntvoltage"`
	Busvoltage   float32 `json:"busvoltage"`
	Loadvoltage  float32 `json:"loadvoltage"`
	CreatedAt    string  `json:"created_at"`
	UpdateAt     string  `json:"update_at"`
}

// BukaDatabase ..
func BukaDatabase() error {
	db, err := gorm.Open("sqlite3", "sukrondb.db")
	if err != nil {
		return err
	}
	db.LogMode(true)
	if db.HasTable(&DataSerialDB{}) == false {
		db.AutoMigrate(&DataSerialDB{})
	}
	KonDB = db
	return nil
}
