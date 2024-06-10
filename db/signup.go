package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/salatielosorno/Gambit/models"
	"github.com/salatielosorno/Gambit/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Start Sign Up")

	err := DbConnect()

	if err != nil {
		return err
	}

	defer Db.Close()

	sentence := fmt.Sprintf("INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('%s','%s', '%s')", sig.UserEmail, sig.UserUUID, tools.MySQLDate())
	fmt.Println(sentence)

	_, err = Db.Exec(sentence)
	if err != nil {
		fmt.Println("An error has occurred::" + err.Error())

		return err
	}

	fmt.Println("Signup > Successfully")
	return nil
}
