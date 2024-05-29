package main

import (
	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Net:                  "tcp",
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		AllowNativePasswords: true,
		ParseTime: true,
	}
	
	sqlStorage := NewMySQLStorage(cfg)
	db, err := sqlStorage.Init()
	if err != nil {
		panic(err)
	}
	store := NewStore(db)
	server := NewAPIServer("0.0.0.0:6000", store)
	server.Serve()
}
