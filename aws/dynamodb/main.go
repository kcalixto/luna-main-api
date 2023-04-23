package dynamodb

type Client struct {
}

func NewDynamoDB() (*Client, error) {
	return &Client{}, nil
}

type PK struct {
	norm string
}

func parsePK() {

}
