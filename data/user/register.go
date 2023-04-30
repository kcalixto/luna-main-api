package data

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	logger "github.com/lunaorg/luna-main-api/libs/log"
	database "github.com/lunaorg/luna-main-api/types/database"
	viewmodel "github.com/lunaorg/luna-main-api/types/viewmodel"
)

func (d *UserData) Register(item viewmodel.RegisterUser) error {
	data, err := database.NewRegisterUserDB(item)
	if err != nil {
		logger.Error("Got error creating NewRegisterUserDB: %s", err.Error())
	}

	av, err := dynamodbattribute.MarshalMap(data)
	if err != nil {
		logger.Error("Got error marshalling new user: %s", err.Error())
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(d.tableName),
	}

	_, err = d.conn.PutItem(input)
	if err != nil {
		logger.Error("Got error calling PutItem: %s", err.Error())
	}

	return nil
}
