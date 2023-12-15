package servers

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"pract7_1/graph/model"
)

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.LogMode(true)
	
	db.Exec(
		"CREATE DATABASE IF NOT EXISTS api_pract_7_database;" +
			"USE api_pract_7_database;")

		db.AutoMigrate(&model.User{}, &model.Product{})
}
