package store

import (
	"fmt"
	"log"

	"github.com/go-pg/pg/v10"
)

func NewPG(conn string) *pg.DB {
	opt, err := pg.ParseURL(conn)
	if err != nil {
		log.Fatal(err)
	}

	db := pg.Connect(opt)

	_, err = db.Exec("SELECT 1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to postgres db")

	return db
}
