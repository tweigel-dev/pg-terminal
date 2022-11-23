package main

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq"
)

var host = os.Getenv("HOST")
var port = os.Getenv("PORT")
var user = os.Getenv("USER")
var password = os.Getenv("PASSWORD")
var dbname = os.Getenv("DATABASE")

var path = os.Getenv("SQL_PATH")


func sendSQL(db *sql.DB, sqlStr string){
	_, e := db.Exec(sqlStr)
	CheckError(e)
}
func readSqlFile(path string) string {
	dat, err := os.ReadFile(path)
	CheckError(err)
	sql_str := string(dat)
	CheckError(err)
	return sql_str
}

func connectDatabase(host string,port string, user string, password string) *sql.DB{
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(fmt.Sprintf("host=%s port=%s user=%s password=<is set> dbname=%s sslmode=disable", host, port, user, dbname))
	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
	return db
}

func main() {
	db := connectDatabase(host,port,user,password)
	sql_string := readSqlFile(path)
	sendSQL(db, sql_string)
	defer db.Close()
}


func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
