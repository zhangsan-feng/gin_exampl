package database

import (
	"admin_backend/global"
	"database/sql"
	"fmt"
	"log"
)

func (Database) createDatabase(databaseName string) Database {
	address := "host=192.168.56.39 user=postgres password=root dbname='' port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	conn, connErr := sql.Open("postgres", address)
	if connErr != nil {
		log.Println("connection postgresql error", connErr)
	}
	if _, sqlErr := conn.Exec(`CREATE DATABASE` + databaseName + `;`); sqlErr != nil {
		log.Println(sqlErr)
	} else {
		log.Println("table create success")
	}
	conn.Close()
	return Database{}
}

func (Database) createTable() {
	TableSql := []string{
		`CREATE TABLE IF NOT EXISTS users(
    		id SERIAL PRIMARY KEY,                
    		login_name VARCHAR(50)  NOT NULL, 
			username VARCHAR(50)  NOT NULL , 
			password VARCHAR(255) NOT NULL,  
			created_time TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_time TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`,

		`CREATE TABLE IF NOT EXISTS role(
    		id SERIAL PRIMARY KEY, 
    		role_name VARCHAR(50) NOT NULL
		);`,

		`CREATE TABLE IF NOT EXISTS power(
			id SERIAL PRIMARY KEY,                
			url_path VARCHAR(100) NOT NULL 
		);`,

		`CREATE TABLE IF NOT EXISTS user_role(
    		user_id INT NOT NULL,
        	role_id INT NOT NULL
		);`,

		`CREATE TABLE IF NOT EXISTS role_power(
    		role_id  INT NOT NULL,
            power_id INT NOT NULL   	
		);`,

		`CREATE TABLE IF NOT EXISTS resources(
    		id SERIAL PRIMARY KEY,
    		categories int NOT NULL,
    		description TEXT,
    		play_count int NOT NULL,
    		created_time TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_time TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`,

		`CREATE TABLE IF NOT EXISTS files(
    		id SERIAL PRIMARY KEY,
    		file_name VARCHAR(100)  NOT NULL,
    		file_type VARCHAR(100)  NOT NULL ,
    		file_size BIGINT NOT NULL,
    		file_md5 VARCHAR(50) NOT NULL,
        	created_time TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_time TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS resource_files(
			resource_id int NOT NULL,
			file_id int NOT NULL
		);`,
		`CREATE TABLE play_record(
    		id SERIAL PRIMARY KEY,
    		resource_id int NOT NULL ,
    		created_time TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_time TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`,
	}
	conn, connErr := sql.Open("postgres", global.PostgresqlAddress)
	if connErr != nil {
		log.Println("connection postgresql error", connErr)
	}

	for i := range TableSql {
		fmt.Println(TableSql[i])
		if _, sqlErr := conn.Exec(TableSql[i]); sqlErr != nil {
			log.Println(sqlErr)
		} else {
			log.Println("table create success")
		}

	}
	conn.Close()
}
