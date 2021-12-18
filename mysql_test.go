package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestMysql(t *testing.T) {

	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8", "root", "root", "127.0.0.1", 3306, "employees")
	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		t.Fatal("failed to connect database ", err)
	}

	create := db.Create(&Employees{EmpNo: 1, FirstName: "wang", LastName: "bear", BirthDate: time.Now(), Gender: "F", HireDate: time.Now()})
	if create.Error != nil {
		t.Fatal(err)
	}

	var emp Employees
	first := db.First(&emp, 1)
	if first.Error != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", emp)
}
