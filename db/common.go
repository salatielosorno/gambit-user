package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/salatielosorno/Gambit/models"
	"github.com/salatielosorno/Gambit/secretm"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	var err error

	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))

	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))

	if err != nil {
		fmt.Println(err.Error())

		return err
	}

	err = Db.Ping()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Successfully connected to the DB!")
	return nil
}

func ConnStr(key models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string

	dbUser = key.Username
	authToken = key.Password
	dbEndpoint = key.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)

	return dsn
}
