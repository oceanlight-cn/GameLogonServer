package main

import (
	"fmt"
	 "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "sync"
)

const (
	_selectUser = "select count(*) from users where name = ? and password = ?"
)

func validate_user(name string, password string) bool{
	var (
		mu sync.Mutex 
    	err error
    )

	mu.Lock()
	defer mu.Unlock()
	
    db,err := sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/game_server")
    if err != nil {
        fmt.Println(err)
        return false
    }
    defer db.Close()

    fmt.Println("db connected.")

    var (
    	user_count int
    )
    err = db.QueryRow(_selectUser,name, password).Scan(&user_count)
    if err != nil {
        fmt.Println(err)
        return false
    }
    //fmt.Println(user_count)
    if user_count == 1 {
    	return true
    } else {
    	return false
    }
}