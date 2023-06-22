package types

import (
	"fmt"

	"github.com/google/uuid"
	viewmodel "github.com/lunaorg/luna-main-api/types/viewmodel"
)

type saveAccountDB struct {
	ID          string `dynamodbav:"account_id"`
	PK          string `dynamodbav:"pk"`
	SK          string `dynamodbav:"sk"`
	AccountName string `dynamodbav:"account_name,omitempty"`
}

func CreateSaveAccountModel(item viewmodel.SaveAccount) (*saveAccountDB, error) {
	UUID := uuid.New().String()

	buildPk := func(id string) string {
		return fmt.Sprintf("%s#%s", "ACCOUNT", id)
	}

	return &saveAccountDB{
		ID:          UUID,
		PK:          buildPk(UUID),
		SK:          "INFO",
		AccountName: item.AccountName,
	}, nil
}

type accountDBTransactionDefault struct {
	ID            string  `dynamodbav:"account_id"`
	PK            string  `dynamodbav:"pk"`
	SK            string  `dynamodbav:"sk"`
	TransactionID string  `dynamodbav:"transaction_id,omitempty"`
	Title         string  `dynamodbav:"title,omitempty"`
	Note          string  `dynamodbav:"note,omitempty"`
	Amount        float64 `dynamodbav:"amount,omitempty"`
}

func CreateAccountTransactionModel(item viewmodel.AccountTransaction, transactionType string) (*accountDBTransactionDefault, error) {
	TRANSACTION_UUID := uuid.New().String()

	buildPk := func(id string) string {
		return fmt.Sprintf("%s#%s#%s", "ACCOUNT", id, "TRANSACTION")
	}

	buildSk := func(id string) string {
		return fmt.Sprintf("%s#%s", transactionType, id)
	}

	return &accountDBTransactionDefault{
		ID:            item.FromAccountID,
		PK:            buildPk(item.FromAccountID),
		SK:            buildSk(TRANSACTION_UUID),
		TransactionID: TRANSACTION_UUID,
		Title:         item.Title,
		Note:          item.Note,
		Amount:        item.Amount,
	}, nil
}
