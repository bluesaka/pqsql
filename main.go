package main

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
)

type User struct {
	tableName struct{} `pg:"user"` // 默认表名是users，这里指定表名为user
	Id        int
	Name      string
	CreatedAt string
	// CreatedAt time.Time
}

func main() {
	Select()
	// Insert()
	// Update()
	// Delete()
}

func Select() {
	// psql -h localhost -p 5432 -U postgres
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "",
		Database: "test",
	})
	defer db.Close()

	user := &User{Id: 11}
	err := db.Model(user).WherePK().Select()
	if err == pg.ErrNoRows {
		fmt.Println("no row")
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("user: %+v\n", user)
}

func Insert() {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "",
		Database: "test",
	})
	defer db.Close()

	user1 := &User{Id: 11, Name: "test1", CreatedAt: time.Now().Format("2006-01-02 15:04:05")}
	result, err := db.Model(user1).Insert()
	if err != nil {
		panic(err)
	}
	fmt.Printf("user: %+v, result: %v\n", user1, result.RowsAffected())
}

func Update() {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "",
		Database: "test",
	})
	defer db.Close()

	user1 := &User{Id: 11, Name: "test2", CreatedAt: time.Now().Format("2006-01-02 15:04:05")}
	result, err := db.Model(user1).WherePK().Update()
	if err != nil {
		panic(err)
	}
	fmt.Printf("user: %+v, result: %v\n", user1, result.RowsAffected())
}

func Delete() {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "",
		Database: "test",
	})
	defer db.Close()

	user1 := &User{Id: 11, Name: "test2", CreatedAt: time.Now().Format("2006-01-02 15:04:05")}
	result, err := db.Model(user1).WherePK().Delete()
	if err != nil {
		panic(err)
	}
	fmt.Printf("user: %+v, result: %v\n", user1, result.RowsAffected())
}
