package data

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/lunaorg/luna-main-api/aws/dynamodb"
)

type UserData struct {
	conn      *dynamodb.DynamoDB
	tableName string
}

func NewUserData() (*UserData, error) {
	_db, err := db.NewDB()
	if err != nil {
		return nil, err
	}

	return &UserData{
		conn:      _db.Svc,
		tableName: _db.TableName,
	}, nil
}
