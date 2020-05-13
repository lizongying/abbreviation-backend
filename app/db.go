package app

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var Db *sql.DB

func InitDb() {
	db, err := sql.Open("sqlite3", "./abbreviation.db")
	if err != nil {
		log.Fatalln(err)
	}
	//_, err = db.Query("CREATE VIRTUAL TABLE IF NOT EXISTS `abbr_v` USING fts5(`id`, `abbr`, `full`, `rate`)")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//_, err = db.Query("INSERT INTO `abbr_v` SELECT * FROM `abbr`")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	Db = db
}
