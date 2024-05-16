package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connection() {
	dsn := "root:12345@tcp(35.202.163.192:3306)/library-system?charset=utf8"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("CONNECTION ERROR", err)
		return
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("PING ERROR", err)
		return
	}

	fmt.Println("CONNECTION OK")
}
