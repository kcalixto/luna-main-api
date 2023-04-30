package types

import (
	"fmt"

	"github.com/google/uuid"
	vm "github.com/lunaorg/luna-main-api/types/viewmodel"
)

type RegisterUserDB struct {
	ID string `dynamodbav:"sk"`
	PK string `dynamodbav:"pk"`
	SK string `dynamodbav:"sk"`

	Name string `dynamodbav:"name,omitempty"`
}

func NewRegisterUserDB(item vm.RegisterUser) (*RegisterUserDB, error) {
	id := uuid.New()

	buildPk := func(login string, id string) string {
		return fmt.Sprintf("%s#%s", login, id)
	}

	return &RegisterUserDB{
		ID:   id.String(),
		PK:   buildPk(item.Login, id.String()),
		SK:   "INFO",
		Name: item.Name,
	}, nil
}
