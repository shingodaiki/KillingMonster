package api

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	// postgres driver
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitStore(dbURL string) {
	var err error
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("cannot open DB conn:", err)
	}
	log.Println("open DB conn OK")

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20) // idle + open
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		log.Fatal("cannot ping DB conn:", err)
	}
	log.Println("ping DB conn OK")

	err = createPixelTable()
	if err != nil {
		log.Fatal("cannot create Pixel table:", err)
	}
	log.Println("create Pixel table OK")

	err = createStatsTable()
	if err != nil {
		log.Fatal("cannot create Stats table:", err)
	}
	log.Println("create Stats table OK")
}

func CloseStore() {
	db.Close()
}

func createPixelTable() error {
	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS pixel(
			id 		serial 			primary key,
			sourcer varchar(32)		not null,
			address varchar(48)		not null,
			browser varchar(128)	not null,
			timeutc varchar(64)		not null)`)
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT column_name, data_type FROM information_schema.columns WHERE table_name='pixel'")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var columnName, dataType string
		err = rows.Scan(&columnName, &dataType)
		if err != nil {
			return err
		}

		switch columnName {
		default:
			return errors.New("unknown columnName")
		case "id":
			if dataType != "integer" {
				return errors.New("id is not an `integer` type")
			}
			break
		case "sourcer":
		case "address":
		case "browser":
		case "timeutc":
			if dataType != "character varying" {
				return errors.New(columnName + " is not a `character varying` type")
			}
			break
		}
	}
	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}

func insertPixel(sourcer, address, browser, timeutc string) {
	_, err := db.Exec(fmt.Sprintf(
		`INSERT INTO pixel(sourcer, address, browser, timeutc)
		VALUES('%s', '%s', '%s', '%s')`,
		sourcer, address, browser, timeutc))
	if err != nil {
		log.Println("cannot store pixel:", err)
	}
}

func createStatsTable() error {
	// todo
	return nil
}

func sumPixelStats() error {
	// todo
	return nil
}
