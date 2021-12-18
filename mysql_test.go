package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestMysql(t *testing.T) {

	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "root", "127.0.0.1", 3306, "employees")
	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		t.Fatal("failed to connect database ", err)
	}

	db.Create(&Employees{EmpNo: 1, FirstName: "wang", LastName: "bear", BirthDate: time.Now(), Gender: "F", HireDate: time.Now()})

	// 读取
	var emp Employees
	db.First(&emp, 1)
	t.Logf("%+v", emp)
}
