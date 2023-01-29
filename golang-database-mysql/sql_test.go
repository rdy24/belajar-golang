package golangdatabase

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

	query := "INSERT INTO customer(id, name) VALUES ('2', 'Tes2')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Insert data success")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime 
		var createdAt time.Time
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating,&birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id)
		fmt.Println("name:", name)
		if email.Valid {
			fmt.Println("email:", email)
		}
		fmt.Println("balance:", balance)
		fmt.Println("rating:", rating)
		fmt.Println("birthDate:", birthDate)
		fmt.Println("married:", married)
		fmt.Println("createdAt:", createdAt)

	}

	defer rows.Close()
}

func TestQuerySqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin '; #"
	password := "admin"

	query := "SELECT username FROM user WHERE username = '"+ username +"' AND password = '"+ password +"'"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("username:", username)
	} else {
		fmt.Println("Username or password is invalid")
	}
}
func TestQuerySqlInjctionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	query := "SELECT username FROM user WHERE username = ? AND password = ?"
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("username:", username)
	} else {
		fmt.Println("Username or password is invalid")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "tes"
	password := "tes"

	query := "INSERT INTO user(username, password) VALUES (?, ?)"
	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Insert data success")
}
func TestExecSqlAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	email := "tees@gmail.com"
	comment := "tesasdjjdsdssada"

	query := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Insert data success", insertId)
}
func TestPrepareStatment(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	

	query := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		email := "tes" + strconv.Itoa(i) + "@gmail.com"
		comment := "tesasdjjdsdssada" + strconv.Itoa(i)
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		lastId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Insert data success", lastId)
	}

	
	fmt.Println("Insert data success")
	defer stmt.Close()
}
func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	query := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	for i := 0; i < 10; i++ {
		email := "tes" + strconv.Itoa(i) + "@gmail.com"
		comment := "tesasdjjdsdssada" + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, query,email, comment)
		if err != nil {
			panic(err)
		}
		lastId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Insert data success", lastId)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}

	
	fmt.Println("Insert data success")
}