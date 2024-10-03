package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const ID_LENGTH = 6

type MySQLDB struct {
	rdatabase *sql.DB
}

func NewMySQLDB(usename string, password string, address string, dbname string) (*MySQLDB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", usename, password, address, dbname))
	if err != nil {
		return nil, fmt.Errorf("couldn't open database: %w", err)
	}

	return &MySQLDB{
		rdatabase: db,
	}, nil
}

func (mdb MySQLDB) SetStatus(requestId string, status string) error {
	query := "UPDATE Requests SET RequestStatus=? WHERE RequestID=?"
	_, err := mdb.rdatabase.Exec(query, status, requestId)
	if err != nil {
		return err
	}
	return nil
}

func (mdb MySQLDB) SetImageCaption(requestId string, imgcap string) error {
	query := "UPDATE Requests SET ImageCaption=? WHERE RequestID=?"
	_, err := mdb.rdatabase.Exec(query, imgcap, requestId)
	if err != nil {
		return err
	}
	return nil
}
