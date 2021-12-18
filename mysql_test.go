package main

import (
	"github.com/jinzhu/gorm"
	"testing"
	"time"
)

func TestMysql(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/employees")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.Create(&Employees{EmpNo: 1, FirstName: "wang", LastName: "bear", BirthDate: time.Now(), Gender: "F", HireDate: time.Now()})

	// 读取
	var emp Employees
	db.First(&emp, 1)
	t.Logf("%+v", emp)
}
