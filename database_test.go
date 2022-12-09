package belajar_go_mysql

import (
	"database/sql"
	"testing"
)
import _ "github.com/go-sql-driver/mysql"

func TestEmpty(t *testing.T) {

}
func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/siakad")
	if err != nil {
		panic(err)
	}
	db.Close()
}

func Test(t *testing.T) {

}
