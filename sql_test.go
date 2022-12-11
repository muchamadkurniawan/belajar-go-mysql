package belajar_go_mysql

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "insert into member (id, name) values ('id1','kurniawan')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new data member")
}

func TestQueryComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "select id, name, email, balance, rating, birth_date, married, create_at from member"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for rows.Next() {
		var id, name, email string
		var balance int32
		var rating float64
		var birthData, createAt time.Time
		var maried bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthData, &maried, &createAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("===================")
		fmt.Println("id : ", id)
		fmt.Println("name :", name)
		fmt.Println("email : ", email)
		fmt.Println("balance :", balance)
		fmt.Println("rating :", rating)
		fmt.Println("birth data :", birthData)
		fmt.Println("married : ", maried)
		fmt.Println("dibuat pada :", createAt)
	}
}
