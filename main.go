package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"
)

import _ "github.com/go-sql-driver/mysql"

// generic query utility
func query(db *sql.DB, expr string, in ...interface{}) ([]interface{}, error) {
	s := "SELECT " + expr + " FROM DUAL"
	log.Println(s)
	rows, err := db.Query(s, in...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.New("No rows!")
	}

	var names []string
	names, err = rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]interface{}, len(names))
	refs := make([]interface{}, len(names))
	for i := range values {
		refs[i] = &(values[i])
	}
	if err = rows.Scan(refs...); err != nil {
		return nil, err
	}
	return values, nil
}

func run(db *sql.DB) error {
	now := time.Now()
	out, err := query(db, "NOW(), ?", now)
	if err != nil {
		return err
	}
	fmt.Printf("column 0: %T %v\n", out[0], out[0])
	fmt.Printf("column 1: %T %v %s\n", out[1], out[1], out[1])

	return nil
}

func main() {
	var dbInfo struct {
		User     string
		Password string
		HostPort string
		Database string
		Timezone string
	}

	flag.StringVar(&dbInfo.User, `user`, "", "database user")
	flag.StringVar(&dbInfo.Password, `password`, "", "database password")
	flag.StringVar(&dbInfo.HostPort, `host`, "", "database host[:port]")
	flag.StringVar(&dbInfo.Database, `database`, "", "database name")
	flag.StringVar(&dbInfo.Timezone, `timezone`, "", "session timezone (@time_zone)")
	flag.Parse()

	connstr := dbInfo.User + ":" + dbInfo.Password + "@tcp(" + dbInfo.HostPort +
		")/" + dbInfo.Database + "?strict=true&charset=utf8&parseTime=True&loc=UTC"
	if dbInfo.Timezone != "" {
		// Quoting with %27 is needed
		// https://github.com/go-sql-driver/mysql/issues/405
		connstr = connstr + "&time_zone=%27" + url.QueryEscape(dbInfo.Timezone) + "%27"
	}
	log.Println(connstr)

	db, err := sql.Open("mysql", connstr)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	if err = run(db); err != nil {
		log.Println(err)
		// As os.Exit shortcircuits defers, we have to close explicitely
		db.Close()
		os.Exit(1)
	}
}
