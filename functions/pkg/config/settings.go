package config

import (
	"os"
)

type MySQLInfo struct {
	User   string
	Pass   string
	Addr   string
	DBName string
}

// DBの設定情報をloadする
func LoadDBConfig() *MySQLInfo {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPass := os.Getenv("MYSQL_PASS")
	mysqlAddr := os.Getenv("MYSQL_ADDR")
	mysqlDBName := os.Getenv("MYSQL_DBNAME")

	config := &MySQLInfo{
		User:   mysqlUser,
		Pass:   mysqlPass,
		Addr:   mysqlAddr,
		DBName: mysqlDBName,
	}

	return config
}
