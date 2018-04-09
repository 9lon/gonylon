package gonylon

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// mssql
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type dbConnectMap map[string]*gorm.DB

var dbConnect dbConnectMap

// InitDB : เชื่อมต่อ DB ทุกตัว
func InitDB(config gin.H) {
	dbConnect = dbConnectMap{}
	//dbConfig := config.GetConfig()["gorm"].(gin.H)
	dbConfig := config["gorm"].(gin.H)
	for key, value := range dbConfig {
		val := value.(gin.H)
		db, _ := gorm.Open(
			val["dialect"].(string),
			"sqlserver://"+val["username"].(string)+":"+val["password"].(string)+"@"+val["host"].(string)+"?database="+val["db"].(string),
		)
		dbConnect[key] = db
	}
}

// GetDB : เลือก DB ที่ต้องการใช้
func GetDB(connectName string) *gorm.DB {
	return dbConnect[connectName]
}

// CloseDB : ปิด DB ทุกตัวที่ทำการ Connect
func CloseDB() {
	for _, dbConn := range dbConnect {
		dbConn.Close()
	}
	fmt.Println("close")
}
