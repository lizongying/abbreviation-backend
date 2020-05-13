package app

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
)

type ReqAbbr struct {
	Keyword string `uri:"keyword" binding:"required"`
}

func GetAbbr(keyword string) string {
	list := make([]string, 0)
	db := Db
	log.Println(keyword)
	rows, err := db.Query("SELECT `abbr`,`full`,`rate` FROM `abbr` WHERE `full` = ? COLLATE NOCASE ORDER BY `rate` LIMIT 10", keyword)
	//rows, err := db.Query("SELECT `abbr`,`full`,`rate` FROM `abbr_v` WHERE `full` MATCH ? ORDER BY `rank` LIMIT 10", keyword)
	if err != nil {
		log.Println("empty")
		return ""
	}
	for rows.Next() {
		var abbr string
		var full string
		var rate int
		if err = rows.Scan(&abbr, &full, &rate); err != nil {
			log.Println(err)
			continue
		}
		log.Println(abbr, full, rate)
		list = append(list, abbr)
	}
	return strings.Join(list, "\n")
}
