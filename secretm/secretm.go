package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/salatielosorno/Gambit/awsgo"
	"github.com/salatielosorno/Gambit/models"
)

func GetSecret(secretName string) (models.SecretRDSJson, error) {
	var dataSecret models.SecretRDSJson

	fmt.Println("Get secret:: " + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		fmt.Println(err.Error())
		return dataSecret, err
	}

	json.Unmarshal([]byte(*key.SecretString), &dataSecret)
	fmt.Println("Read secret successfully!")

	return dataSecret, nil
}
