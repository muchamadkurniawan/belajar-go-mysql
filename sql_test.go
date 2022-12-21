package belajar_go_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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
		var id, name string
		var email sql.NullString
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
		if email.Valid {
			fmt.Println("email : ", email.String)
		}
		fmt.Println("balance :", balance)
		fmt.Println("rating :", rating)
		fmt.Println("birth data :", birthData)
		fmt.Println("married : ", maried)
		fmt.Println("dibuat pada :", createAt)
	}
}

func TestSQLinjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "kurniawan"

	query := "select username from user where username = '" + username +
		"' and passwod = '" + password + "'limit 1"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses login ", username)
	} else {
		fmt.Println("gagal login")
	}
}

func TestSQLinjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "kurniawan"

	query := "select username from user where username = ? and passwod = ? limit 1"
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses login ", username)
	} else {
		fmt.Println("gagal login")
	}
}

func TestExecSqlSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	id := "member4"
	name := "kurkur"
	ctx := context.Background()
	query := "insert into member (id, name) values (?,?)"
	_, err := db.ExecContext(ctx, query, id, name)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new data member")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	email := "kurniawan@gmail.com"
	comment := "hallo ini comment pertama"

	query := "insert into comment(email, comment) values (?,?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("success insert new comment by", insertID)
}

func TestPrepareStetment(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	//email := "kurniawan@gmail.com"
	//comment := "hallo ini comment kedua"

	query := "insert into comment(email, comment) values (?,?)"

	stetment, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(nil)
	}

	for i := 1; i <= 10; i++ {
		//fmt.Println("iiiii :", i)
		email := "kurniawan" + strconv.Itoa(i) + "@gmail.com"
		comment := "commentar ke-" + strconv.Itoa(i)
		result, err := stetment.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("last id adalah ", id)

	}

	defer stetment.Close()
}

func TestSQLTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(nil)
	}

	//do transaction
	query := "insert into comment(email, comment) values (?,?)"
	for i := 1; i <= 10; i++ {
		//fmt.Println("iiiii :", i)
		email := "kurniawan" + strconv.Itoa(i) + "@gmail.com"
		comment := "commentar ke-" + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, query, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("last id adalah ", id)

	}
	//err = tx.Rollback() // untuk rollback
	err = tx.Commit() //untuk commit
	if err != nil {
		panic(nil)
	}
}
