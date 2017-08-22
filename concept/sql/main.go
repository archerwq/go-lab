// See http://go-database-sql.org/overview.html

package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	createTableSQL = `
		CREATE TABLE IF NOT EXISTS employee (
  			id int(6) unsigned NOT NULL AUTO_INCREMENT,
			name varchar(30) NOT NULL,
			city varchar(30),
			PRIMARY KEY(id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8
	`
)

type employee struct {
	ID   int64
	Name string
	City string
}

var employeeMap = map[string]string{
	"Qiang Wang":  "Beijing",
	"Sen Song":    "Beijing",
	"Yuxin Jiang": "Beijing",
	"Wanli Xiao":  "Melbourne",
	"Zhiyu Liu":   "Beijing",
	"Peng Liu":    "Beijing",
}

func main() {
	// To access databases in Go, you use a sql.DB. You use this type to create
	// statements and transactions, execute queries, and fetch results.
	// sql.DB performs some important tasks for you behind the scenes:
	// It opens and closes connections to the actual underlying database, via the driver.
	// It manages a pool of connections as needed, which may be a variety of things as mentioned.
	// The sql.DB abstraction is designed to keep you from worrying about how to manage concurrent access
	// to the underlying datastore. A connection is marked in-use when you use it to perform a task,
	// and then returned to the available pool when it’s not in use anymore. One consequence of this is that
	// if you fail to release connections back to the pool, you can cause sql.DB to open a lot of connections,
	// potentially running out of resources (too many connections, too many open file handles,
	// lack of available network ports, etc).
	db, err := sql.Open("mysql", "root:test@tcp(127.0.0.1:4306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// ping
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	// drop table if exists
	if _, err := db.ExecContext(ctx, "DROP TABLE IF EXISTS employee;"); err != nil {
		log.Fatal(err)
	}

	// create table
	if _, err := db.ExecContext(ctx, createTableSQL); err != nil {
		log.Fatal(err)
	}

	// insert
	stmt, err := db.PrepareContext(ctx, "INSERT INTO employee(name,city) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for name, city := range employeeMap {
		res, err := stmt.ExecContext(ctx, name, city)
		if err != nil {
			log.Fatal(err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		count, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID = %d, affected = %d\n", id, count)
	}

	// query row
	var (
		id   int64
		name string
		city string
	)
	err = db.QueryRowContext(ctx, "SELECT * FROM employee WHERE NAME=?", "Qiang Wan").Scan(&id, &name, &city)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("not found: [%s]\n", "Qiang Wang")
	case err != nil:
		log.Fatal(err)
	default:
		log.Printf("querying [%s] got: ID = %d, name = %s, city = %s\n", "Qiang Wang", id, name, city)
	}

	// query rows
	rows, err := db.QueryContext(ctx, "SELECT * FROM employee")
	if err != nil {
		log.Fatal(err)
	}
	// You should always defer rows.Close(), even if you also call rows.Close()
	// explicitly at the end of the loop, which isn’t a bad idea.
	// rows.Close() is a harmless no-op if it’s already closed, so you can call it multiple times.
	// Notice, however, that we check the error first, and only call rows.Close() if there isn’t an error,
	// in order to avoid a runtime panic.
	defer rows.Close()
	// As long as there’s an open result set (represented by rows), the underlying connection
	// is busy and can’t be used for any other query. That means it’s not available in the connection pool.
	// If you iterate over all of the rows with rows.Next(), eventually you’ll read the last row,
	// and rows.Next() will encounter an internal EOF error and call rows.Close() for you.
	// But if for some reason you exit that loop – an early return, or so on –
	// then the rows doesn’t get closed, and the connection remains open.
	// (It is auto-closed if rows.Next() returns false due to an error, though).
	//  This is an easy way to run out of resources.
	for rows.Next() {
		// When you iterate over rows and scan them into destination variables,
		// Go performs data type conversions work for you, behind the scenes.
		// It is based on the type of the destination variable.
		if err := rows.Scan(&id, &name, &city); err != nil {
			log.Fatal(err)
		}
		log.Printf("ID = %d, name = %s, city = %s\n", id, name, city)
	}
	if err := rows.Close(); err != nil {
		log.Fatal(err)
	}
	// The error from rows.Err() could be the result of a variety of errors in the rows.Next() loop.
	// The loop might exit for some reason other than finishing the loop normally,
	// so you always need to check whether the loop terminated normally or not.
	// An abnormal termination automatically calls rows.Close(), although it’s harmless to call it multiple times.
	// You should always check for an error at the end of the for rows.Next() loop.
	// If there’s an error during the loop, you need to know about it.
	// Don’t just assume that the loop iterates until you’ve processed all the rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// transaction
	// In Go, a transaction is essentially an object that reserves a connection to the datastore.
	// It lets you do all of the operations we’ve seen thus far, but guarantees that they’ll be executed
	// on the same connection. You begin a transaction with a call to db.Begin(), and close it with a Commit()
	// or Rollback() method on the resulting Tx variable. Under the covers, the Tx gets a connection from the pool,
	// and reserves it for use only with that transaction. The methods on the Tx map one-for-one to methods
	// you can call on the database itself, such as Query() and so forth.
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.ExecContext(ctx, "UPDATE employee SET city=? WHERE name=?", "Zhuhai", "Qiang Wang")
	if err != nil {
		if err := tx.Rollback(); err != nil {
			log.Fatal(err)
		}
		log.Fatal(err)
	}
	_, err = tx.ExecContext(ctx, "UPDATE employee SET city=? WHERE name=?", "Guangzhou", "Yuxin Jiang")
	if err != nil {
		if err := tx.Rollback(); err != nil {
			log.Fatal(err)
		}
		log.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
