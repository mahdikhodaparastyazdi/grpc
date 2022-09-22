package service

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Service struct {
	db *sql.DB
}

func NewService() *Service {

	return &Service{db: connection("mysql", "root:@/user", "")}

}

func connection(driverName string, userName string, password string) *sql.DB {
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open(driverName, userName+":"+password+"@(127.0.0.1:80)/user")
	//fmt.Println(driverName, userName+":"+password+"@(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	} else {
		return db
	}
}
func (h *Service) Add(ctx context.Context, name string) bool {
	test := connection("mysql", "root:@/user", "")
	fmt.Println(test)
	q := fmt.Sprintf("insert into orders (name) VALUES ('%s')", name)
	fmt.Println(q)
	_, err := h.db.ExecContext(ctx, q)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("hereee")
	return true
}
