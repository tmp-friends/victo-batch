package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// 環境変数をloadする
func LoadEnv() map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := map[string]string{
		"BEARER_TOKEN": os.Getenv("BEARER_TOKEN"),
	}

	return env
}

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
