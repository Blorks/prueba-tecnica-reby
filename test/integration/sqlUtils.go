package integration

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gorm.io/gorm"
)

func LoadSQLFile(dbConn *gorm.DB, sqlFile string) {
	file, err := ioutil.ReadFile(sqlFile)

	if err != nil {
		fmt.Println(err)
	}

	tx := dbConn.Begin()

	defer func() {
		tx.Rollback()
	}()

	for _, q := range strings.Split(string(file), ";") {
		q := strings.TrimSpace(q)

		if q == "" {
			continue
		}

		tx.Exec(q)
	}

	tx.Commit()
}
