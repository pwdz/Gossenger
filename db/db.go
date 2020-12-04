package db

import (
	"database/sql"
	"os"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"Gossenger/constants"
)

type DB struct{
	db *sql.DB
}

func CreateDB() *DB {
	if _, err := os.Stat(constants.DBBasePath); os.IsNotExist(err) {
		os.MkdirAll(constants.DBBasePath, 0755)
	}

	database, _ := sql.Open("sqlite3", constants.DBBasePath+constants.DBName)
    statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users (username VARCHAR(50), password VARCHAR(50), primary key (username) )")
	statement.Exec()
	
	fmt.Println("[#] DB init.")
	return &DB{
		db: database,
	}
}
func (db *DB) AddData(username,password string){
	fmt.Println("[#] DB Add data:",username,password)

	database := db.db
	statement, _ := database.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
	_, err := statement.Exec(username, password)
	if err != nil{
		fmt.Println("[#] DB problem adding:", err.Error())
	}
}
func (db *DB) ChangeUsername(oldname, newname string){
	database := db.db

	statement,_ := database.Prepare("UPDATE users SET username = ? WHERE username = ?");
	statement.Exec(newname, oldname)
}
func (db *DB) GetPassword(username string)string{
	database := db.db
	statement,_ := database.Prepare("SELECT password FROM users WHERE username = ?");
	rows, _ := statement.Query(username)
	
	var password string
	for rows.Next(){
		rows.Scan(&password)
		return password
	}
	return ""
}
func (db *DB) DoesExist(username string)bool{
	fmt.Println("[#] DB Does Exist func.")



	database := db.db
	statement,_ := database.Prepare("SELECT username FROM users WHERE username = ?");
	rows, _ := statement.Query(username)
	
	for rows.Next(){
		return true
	}
	return false
}
