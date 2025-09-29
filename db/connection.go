package db

import (
	"fmt"

	"github.com/enghasib/laundry_service/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func getConnectionString(cnf *config.Config) string {
	cnnSource := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		cnf.DBUserName, cnf.DBPassword, cnf.DBHost, cnf.DBPort, cnf.DBName,
	)
	fmt.Println("connection string", cnnSource)
	return cnnSource
	// return "user= password=postgres host=localhost port=5433 dbname=laundry_service sslmode=disable"
}

func NewConnection(cnf *config.Config) (*sqlx.DB, error) {
	dbSource := getConnectionString(cnf)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println("DB err:", err)
		return nil, err
	}

	return dbCon, nil

}
