package config

type dbConfig struct {
	MigrationsFilePath string
}

var  DBConfig = dbConfig{
	MigrationsFilePath: "/Users/klook/workspace/src/github.com/ItsWewin/the-general-public/database/migrations",
}

