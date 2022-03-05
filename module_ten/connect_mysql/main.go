package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hjtcn/benbird-and-annie-go/pkg/apis"
	"log"
)

func main() {
	db_obj, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_learning")

	if err != nil {
		log.Fatal(err)
	}

	defer db_obj.Close()
	if err := db_obj.Ping(); nil != err {
		fmt.Println("DB 连接测试失败")
	}
	rows, _ := db_obj.Query("select name, age, sex from personal_info")
	//persons := []*pkg.PersonalInfo{}
	p := &apis.PersonalInfoList{}

	for rows.Next() {
		var name string
		var age int
		var sex string

		rows.Scan(&name, &age, &sex)

		p.Items = append(p.Items, &apis.PersonalInfo{
			Name: name,
			Age: int64(age),
			Sex: sex,
		})
		//fmt.Printf("name: %s\t age: %d\n", name, age)
	}

	data, _ := json.Marshal(p)

	fmt.Printf("%s", data)
}
