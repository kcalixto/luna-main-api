package aws

type AWS interface {
	DynamoDBClient() string
}
