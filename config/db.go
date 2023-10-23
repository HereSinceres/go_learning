package config

import (
	"fmt"
)

const (
	DBUser     = "user"
	DBPassword = "user"
	DBName     = "mydb"
	DBHost     = "localhost"
	DBPort     = "3306"
	DBtype     = "mysql"
)

func GetDBType() string {
	return DBtype
}

func GetMySQLConnectionString() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser,
		DBPassword,
		DBHost,
		DBPort,
		DBName)

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return dsn
}
