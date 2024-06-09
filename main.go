package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/salatielosorno/Gambit/awsgo"
	"github.com/salatielosorno/Gambit/db"
	"github.com/salatielosorno/Gambit/models"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitializeAWS()

	if !ValidParameters() {
		fmt.Println("Error in the parameters. You must send 'SecretName'")
		err := errors.New("Error in the parameters. You must send 'SecretName'")
		return event, err
	}

	var data models.SignUp

	for labelRow, valueAttr := range event.Request.UserAttributes {
		switch labelRow {
		case "email":
			data.UserEmail = valueAttr
			fmt.Println("Email::" + data.UserEmail)
		case "sub":
			data.UserUUID = valueAttr
			fmt.Println("Sub::" + data.UserUUID)
		}
	}

	err := db.ReadSecret()

	if err != nil {
		fmt.Println("Error reading secret:: " + err.Error())
		return event, err
	}

	err = db.SignUp(data)

	return event, err
}

func ValidParameters() bool {
	var paramExist bool

	_, paramExist = os.LookupEnv("SecretName")

	return paramExist
}
