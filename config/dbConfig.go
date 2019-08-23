package config

type dbConfig struct {
	MigrationsFilePath string
	Host string
	Port uint16
	DBName string
	UserName string
	PassWord string
	Protocol string
}

var  DBConfig = dbConfig{
	MigrationsFilePath: "/Users/klook/workspace/src/github.com/ItsWewin/the-general-public/database/migrations",
	Host: "127.0.0.1",
	Port: 3306,
	DBName: "general_public",
	UserName: "root",
	PassWord: "wewin123",
	Protocol: "tcp",
}

