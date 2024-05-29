package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MySQLStorage struct{
	db *sql.DB
}

func NewMySQLStorage(cfg mysql.Config) *MySQLStorage{
	db, err := sql.Open("mysql", cfg.FormatDSN());
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping();
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MYSQL")
	return &MySQLStorage{db}
}

func (s *MySQLStorage) Init() (*sql.DB, error){
	err := s.createUsersTable()
	return s.db, err
}

func (s * MySQLStorage) createUsersTable() error {
	_, err := s.db.Exec(`
			CREATE TABLE IF NOT EXISTS users (
				id INT UNSIGNED NOT NULL AUTO_INCREMENT,
				email VARCHAR(255) NOT NULL,
				firstName VARCHAR(255) NOT NULL,
				lastName VARCHAR(255) NOT NULL,
				password VARCHAR(255) NOT NULL,
				createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				PRIMARY KEY (id),
				UNIQUE KEY (email)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)
	if err != nil {
		return err
	}
	return nil
}

