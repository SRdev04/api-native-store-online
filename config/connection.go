package config

import (
	"database/sql"
	"time"

	"github.com/SRdev04/api-native-store-online/helper"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionSql() *sql.DB {
	db, err := sql.Open("mysql", "root:sahrulramadhan@tcp(localhost:3306)/belajar_restapi?parseTime=true")
	helper.IfError(err)

	db.SetMaxIdleConns(30)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db

}
