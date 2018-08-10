package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func printaccounts() {
	db, err := sql.Open("mysql", "user:qburstasd@tcp(localhost:3306)/mydb")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Connection Established")
	}
	name := ""
	password := ""
	rows, err := db.Query("SELECT * FROM ACCOUNTS;")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name, &password)
		if err != nil {
			log.Println(err)
		}
		log.Println(name, password)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
}

func insert(user, password string) (message string) {
	db, err := sql.Open("mysql", "user:qburstasd@tcp(localhost:3306)/mydb")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Connection Established")
	}

	// CREATES TABLE IF THERE IS NO TABLE
	stmt0, err := db.Prepare("CREATE TABLE IF NOT EXISTS ACCOUNTS(NAME VARCHAR(30) PRIMARY KEY,PASSWORD VARCHAR(20))")
	if err != nil {
		log.Println(err)
		message = "negative"
	}
	res, err := stmt0.Exec(user, password)

	stmt, err := db.Prepare("INSERT INTO ACCOUNTS VALUES(?,?)")
	if err != nil {
		fmt.Println("HERE0")
		log.Println(err)
		message = "negative"
	}
	res, err = stmt.Exec(user, password)
	if err != nil {
		fmt.Println("HERE1")
		log.Println(err)
		message = "negative"
		fmt.Println("HERE1F")
	} else {
		lastID, err := res.LastInsertId()
		fmt.Println("HERE2")
		if err != nil {
			fmt.Println("HERE2")
			log.Println(err)
			message = "negative"
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			fmt.Println("HERE3")
			log.Println(err)
			message = "negative"
		}
		log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
	}
	defer db.Close()
	return message
}

func checkuser(cuser string, cpassword string) string {
	result := "no"
	db, err := sql.Open("mysql", "user:qburstasd@tcp(localhost:3306)/mydb")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Connection Established")
	}
	name := ""
	password := ""
	rows, err := db.Query("SELECT * FROM ACCOUNTS WHERE NAME= '" + cuser + "';")
	if err != nil {
		log.Println(err)
		fmt.Println("HERE0")
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&name, &password)
		if err != nil {
			log.Println(err)
		}
		log.Println(name, password, cuser, cpassword)
		if name == cuser && password == cpassword {
			result = "yes"
		}
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		fmt.Println("HERE2")
	}
	defer db.Close()

	return result
}
