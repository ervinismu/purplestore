package db

import (
	"fmt"

	// "github.com/RichardKnop/machinery/v1/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB(dbDriver string, dbConnect string) (*sqlx.DB, error) {
	db, err := sqlx.Open(dbDriver, dbConnect)
	if err != nil {
		errMsg := fmt.Errorf("database error connect : %w", err)
		return nil, errMsg
	}

	err = db.Ping()
	if err != nil {
		errMsg := fmt.Errorf("database error ping : %w", err)
		return nil, errMsg
	}

	fmt.Println("success connect to database")

	return db, nil
}
