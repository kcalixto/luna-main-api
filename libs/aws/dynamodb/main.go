package db

import (
	logger "github.com/lunaorg/luna-main-api/libs/log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DB struct {
	Svc       *dynamodb.DynamoDB
	TableName string
}

func NewDB() (*DB, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(endpoints.SaEast1RegionID),
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	svc := dynamodb.New(sess)

	return &DB{
		Svc:       svc,
		TableName: os.Getenv("LUNA_INFRA_DB_NAME"),
	}, nil
}
